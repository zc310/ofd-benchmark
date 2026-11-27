// ofd2png-benchmark measures native OFD -> PNG converters at a fixed 200 DPI.
package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/zc310/ofd-benchmark/internal/benchmarkdata"
	"github.com/zc310/ofd-benchmark/internal/converter"
)

type pngConverter struct {
	name    string
	command []string
	tables  string
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: ofd2png-benchmark <input.ofd> [converter ...] (200 DPI; see docs/ofd-to-png.md)")
	}
	input, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}
	if _, err := os.Stat(input); err != nil {
		return err
	}
	name := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))
	if name != "999" && name != "intro" && name != "zsbk" {
		return fmt.Errorf("no 200 DPI PNG benchmark configured for %q (use 999, intro, zsbk)", name)
	}
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	bin := filepath.Dir(exe)
	base := filepath.Dir(bin)
	refPages, _, err := inspectPages(filepath.Join(base, "testdata", "suwell", name), name)
	if err != nil {
		return fmt.Errorf("reference PNGs: %w", err)
	}
	timeout, err := converter.Timeout()
	if err != nil {
		return err
	}
	converters := pngConverters(bin, base)
	selected := args[1:]
	if len(selected) == 0 {
		selected = []string{"zc310", "ofdgo", "ofdrw-rs", "ofdbox-folio", "sofd",
			"java", "ofdmanager", "ofd-utility", "linuxdeepin-rofd"}
	}
	for _, key := range selected {
		if _, ok := converters[key]; !ok {
			return fmt.Errorf("unknown PNG converter %q", key)
		}
	}
	outputRoot := filepath.Join(converter.OutputRoot(), "output_png")
	if err := os.MkdirAll(outputRoot, 0o755); err != nil {
		return err
	}
	var updates []benchmarkdata.CellUpdate
	for _, key := range selected {
		c := converters[key]
		values := map[string]string{}
		// Reset quality as well as speed/size; never reuse a previous run's score.
		for _, metric := range []string{"speed", "cpu", "memory", "ms-per-page", "size",
			"pixel-similarity", "content-similarity", "foreground-f1", "ssim"} {
			values[c.tables+"-"+metric] = "FAILED"
		}
		destination := filepath.Join(outputRoot, name+"_"+c.name)
		if err := os.RemoveAll(destination); err != nil {
			return err
		}
		scratch, err := os.MkdirTemp(outputRoot, ".png-run-")
		if err != nil {
			return err
		}
		output := filepath.Join(scratch, "pages")
		command := append(append([]string{}, c.command...), input, output)
		duration, cpu, memory, runErr := converter.RunDirectory(c.name, command, output, timeout)
		if runErr == nil {
			pages, size, err := inspectPages(output, "page")
			runErr = err
			if runErr == nil && pages != refPages {
				runErr = fmt.Errorf("page count mismatch: reference=%d, output=%d", refPages, pages)
			}
			if runErr == nil {
				runErr = os.Rename(output, destination)
			}
			if runErr == nil {
				values[c.tables+"-speed"] = benchmarkdata.FormatDuration(duration)
				values[c.tables+"-cpu"] = benchmarkdata.FormatDuration(cpu)
				values[c.tables+"-memory"] = converter.FormatPeakRSS(memory)
				values[c.tables+"-ms-per-page"] = fmt.Sprintf("%.2fms", float64(duration.Microseconds())/1000/float64(pages))
				values[c.tables+"-size"] = benchmarkdata.FormatPDFSize(size)
				fmt.Printf("Path: %s, time: %v, CPU: %v, peak RSS: %s, pages: %d, PNG total: %s\nSaved: %s\n",
					c.tables, duration, cpu, converter.FormatPeakRSS(memory), pages, values[c.tables+"-size"], destination)
			}
		}
		if err := os.RemoveAll(scratch); err != nil {
			return err
		}
		if runErr != nil {
			fmt.Printf("FAILED: %v\n", runErr)
		}
		for table, value := range values {
			updates = append(updates, benchmarkdata.CellUpdate{
				Report: "ofd-to-png", Table: table, Input: name + ".ofd", Converter: pngRowName(c.name), Value: value,
			})
		}
	}
	return benchmarkdata.Update(filepath.Join(base, "docs", "benchmark-results.yaml"), updates, false)
}

func pngRowName(name string) string {
	if label, ok := map[string]string{
		"go-zc310":  "go-zc310-ofd",
		"go-ofdgo":  "go-xiaoqidun-ofdgo",
		"rust-sofd": "rust-sOFD",
	}[name]; ok {
		return label
	}
	return name
}

func pngConverters(bin, base string) map[string]pngConverter {
	return map[string]pngConverter{
		"zc310":            {"go-zc310", []string{filepath.Join(bin, "zc310-png")}, "png"},
		"ofdgo":            {"go-ofdgo", []string{filepath.Join(bin, "ofdgo-png")}, "png"},
		"ofd-utility":      {"rust-ofd-utility", []string{"python3", filepath.Join(base, "tools", "ofd_utility_png.py")}, "png"},
		"ofdrw-rs":         {"go-ofdrw-rs", []string{filepath.Join(bin, "ofdrw-rs-png")}, "png"},
		"ofdbox-folio":     {"rust-ofdbox-folio", []string{filepath.Join(bin, "ofdbox-folio-convert"), "--png"}, "png"},
		"sofd":             {"rust-sofd", []string{"python3", filepath.Join(base, "tools", "native_cli_png.py"), "sofd"}, "png"},
		"java":             {"java-ofdrw", []string{"bash", filepath.Join(base, "java", "ofd2png_java.sh")}, "png"},
		"ofdmanager":       {"rust-ofdmanager", []string{"python3", filepath.Join(base, "tools", "native_cli_png.py"), "ofdmanager"}, "png"},
		"linuxdeepin-rofd": {"rust-linuxdeepin-rofd", []string{filepath.Join(bin, "linuxdeepin-rofd-png")}, "png"},
	}
}

// inspectPages checks PNG headers and a complete 1-based page set, not lexical
// filename order (page_10 must follow page_9, not page_1).
func inspectPages(dir, prefix string) (int, int64, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, 0, err
	}
	if len(entries) == 0 {
		return 0, 0, fmt.Errorf("no PNG pages in %s", dir)
	}
	var size int64
	for _, entry := range entries {
		number := strings.TrimSuffix(strings.TrimPrefix(entry.Name(), prefix+"_"), ".png")
		page, err := strconv.Atoi(number)
		if err != nil || page < 1 || page > len(entries) || entry.Name() != fmt.Sprintf("%s_%d.png", prefix, page) || !entry.Type().IsRegular() {
			return 0, 0, fmt.Errorf("unexpected or nonconsecutive page: %s", entry.Name())
		}
		file, err := os.Open(filepath.Join(dir, entry.Name()))
		if err != nil {
			return 0, 0, err
		}
		_, decodeErr := png.DecodeConfig(file)
		closeErr := file.Close()
		if decodeErr != nil {
			return 0, 0, fmt.Errorf("%s: %w", entry.Name(), decodeErr)
		}
		if closeErr != nil {
			return 0, 0, closeErr
		}
		info, err := entry.Info()
		if err != nil {
			return 0, 0, err
		}
		size += info.Size()
	}
	return len(entries), size, nil
}
