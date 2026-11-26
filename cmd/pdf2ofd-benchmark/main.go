package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/zc310/ofd-benchmark/internal/benchmarkdata"
	"github.com/zc310/ofd-benchmark/internal/converter"
)

type Converter struct {
	Name       string
	Command    []string
	OutputFile string
}

func pdf2ofdPython() string {
	if path := os.Getenv("PDF2OFD_PYTHON"); path != "" {
		return path
	}
	return "python3"
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: pdf2ofd-benchmark <file.pdf> [converter1] [converter2] ...")
	}
	timeout, err := converter.Timeout()
	if err != nil {
		log.Fatal(err)
	}

	inputFile, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	binDir := filepath.Join(filepath.Dir(exePath))

	// Create output directory
	outputDir := filepath.Join(binDir, "output_ofd")
	os.MkdirAll(outputDir, 0755)

	// Get input filename without extension
	inputBase := filepath.Base(inputFile)
	inputName := inputBase[:len(inputBase)-len(filepath.Ext(inputBase))]

	converters := map[string]Converter{
		"zc310": {
			Name:       "pdf2ofd-zc310",
			Command:    []string{filepath.Join(binDir, "ofd-converter"), "--from", "pdf", "-o", filepath.Join(binDir, "output_zc310.ofd"), inputFile},
			OutputFile: filepath.Join(binDir, "output_zc310.ofd"),
		},
		"ofdgo": {
			Name:       "pdf2ofd-ofdgo",
			Command:    []string{filepath.Join(binDir, "ofdgo-pdf2ofd"), inputFile, filepath.Join(binDir, "output_ofdgo.ofd")},
			OutputFile: filepath.Join(binDir, "output_ofdgo.ofd"),
		},
		"java": {
			Name:       "pdf2ofd-ofdrw",
			Command:    []string{filepath.Join(binDir, "..", "java", "pdf2ofd_java.sh"), inputFile, filepath.Join(binDir, "output_java.ofd")},
			OutputFile: filepath.Join(binDir, "output_java.ofd"),
		},
		"pdf2ofd": {
			Name:       "python-pdf2ofd",
			Command:    []string{pdf2ofdPython(), filepath.Join(binDir, "..", "python", "pdf2ofd_convert.py"), inputFile, filepath.Join(binDir, "output_pdf2ofd.ofd")},
			OutputFile: filepath.Join(binDir, "output_pdf2ofd.ofd"),
		},
		"jsofd": {
			Name:       "node-jsofd",
			Command:    []string{"node", filepath.Join(binDir, "..", "node", "pdf2ofd_jsofd.mjs"), inputFile, filepath.Join(binDir, "output_jsofd.ofd")},
			OutputFile: filepath.Join(binDir, "output_jsofd.ofd"),
		},
	}

	selectedConverters := os.Args[2:]
	if len(selectedConverters) == 0 {
		selectedConverters = []string{"zc310", "ofdgo", "java", "pdf2ofd", "jsofd"}
	}

	fmt.Printf("Input file: %s\n\n", inputFile)

	type result struct {
		duration time.Duration
		cpu      time.Duration
		memory   int64
		size     int64
		ok       bool
	}
	results := make(map[string]result)

	for _, name := range selectedConverters {
		if c, ok := converters[name]; ok {
			duration, cpu, memory, size, ok := converter.Run(c.Name, c.Command, c.OutputFile, timeout)
			results[c.Name] = result{duration, cpu, memory, size, ok}

			if ok {
				// Move to output directory
				dstFile := filepath.Join(outputDir, fmt.Sprintf("%s_%s%s", inputName, c.Name, filepath.Ext(c.OutputFile)))
				if err := os.Rename(c.OutputFile, dstFile); err != nil {
					fmt.Printf("Failed to save: %v\n", err)
				} else {
					fmt.Printf("Saved: %s\n", filepath.Base(dstFile))
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("Unknown converter: %s\n\n", name)
		}
	}

	fmt.Println("=== Summary ===")
	for _, name := range selectedConverters {
		if c, ok := converters[name]; ok {
			if r, ok := results[c.Name]; ok {
				if r.ok {
					sizeStr := fmt.Sprintf("%.2f KB", float64(r.size)/1024)
					fmt.Printf("%-20s %10v  CPU: %v  peak RSS: %s  %s\n", c.Name+":", r.duration, r.cpu, converter.FormatPeakRSS(r.memory), sizeStr)
				} else {
					fmt.Printf("%-20s %10v  FAILED\n", c.Name+":", r.duration)
				}
			}
		}
	}

	projectDir := filepath.Dir(binDir)
	var refSize int64
	if info, err := os.Stat(inputFile); err == nil {
		refSize = info.Size()
	}
	rowNames := map[string]string{
		"zc310":   "go-zc310-ofd",
		"ofdgo":   "go-xiaoqidun-ofdgo",
		"java":    "java-ofdrw",
		"pdf2ofd": "python-pdf2ofd",
		"jsofd":   "node-jsofd",
	}
	var updates []benchmarkdata.CellUpdate
	for _, name := range selectedConverters {
		c, exists := converters[name]
		if !exists {
			continue
		}
		r, exists := results[c.Name]
		rowName, supported := rowNames[name]
		if !exists || !supported {
			continue
		}
		elapsed, cpu, memory, size, ratio := "FAILED", "FAILED", "FAILED", "FAILED", "FAILED"
		if r.ok {
			elapsed = benchmarkdata.FormatDuration(r.duration)
			cpu = benchmarkdata.FormatDuration(r.cpu)
			memory = converter.FormatPeakRSS(r.memory)
			size = benchmarkdata.FormatOFDSize(r.size)
			ratio = benchmarkdata.FormatRatio(r.size, refSize)
		}
		for table, value := range map[string]string{
			"pdf-speed":      elapsed,
			"pdf-cpu":        cpu,
			"pdf-memory":     memory,
			"pdf-ofd-size":   size,
			"pdf-size-ratio": ratio,
		} {
			updates = append(updates, benchmarkdata.CellUpdate{
				Report: "pdf-to-ofd", Table: table, Input: inputBase,
				Converter: rowName, Value: value,
			})
		}
	}
	if len(updates) > 0 {
		if err := benchmarkdata.Update(filepath.Join(projectDir, "docs", "benchmark-results.yaml"), updates, false); err != nil {
			log.Fatalf("update benchmark YAML: %v", err)
		}
		fmt.Println("Updated docs/benchmark-results.yaml")
	}
}
