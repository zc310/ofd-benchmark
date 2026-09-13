package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Converter struct {
	Name       string
	Command    []string
	OutputFile string
}

func runConverter(name string, args []string, outputFile string) (time.Duration, int64) {
	fmt.Printf("=== %s ===\n", name)
	start := time.Now()

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "/"

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	duration := time.Since(start)

	var size int64
	if info, err := os.Stat(outputFile); err == nil {
		size = info.Size()
	}

	fmt.Printf("Time: %v\n", duration)
	return duration, size
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ofd-benchmark <file.ofd> [converter1] [converter2] ...")
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

	converters := map[string]Converter{
		"ofdgo": {
			Name:       "github.com/xiaoqidun/ofdgo",
			Command:    []string{filepath.Join(binDir, "ofdgo-convert"), inputFile, filepath.Join(binDir, "output_ofdgo.pdf")},
			OutputFile: filepath.Join(binDir, "output_ofdgo.pdf"),
		},
		"zc310": {
			Name:       "github.com/zc310/ofd",
			Command:    []string{filepath.Join(binDir, "zc310-convert"), inputFile, filepath.Join(binDir, "output_zc310.pdf")},
			OutputFile: filepath.Join(binDir, "output_zc310.pdf"),
		},
		"rust": {
			Name:       "github.com/easy-4-rust/easyofd-rust",
			Command:    []string{filepath.Join(binDir, "easyofd"), "to-pdf", inputFile, filepath.Join(binDir, "output_rust.pdf")},
			OutputFile: filepath.Join(binDir, "output_rust.pdf"),
		},
	}

	selectedConverters := os.Args[2:]
	if len(selectedConverters) == 0 {
		selectedConverters = []string{"ofdgo", "zc310", "rust"}
	}

	fmt.Printf("Input file: %s\n\n", inputFile)

	results := make(map[string]time.Duration)
	sizes := make(map[string]int64)

	for _, name := range selectedConverters {
		if c, ok := converters[name]; ok {
			duration, size := runConverter(c.Name, c.Command, c.OutputFile)
			results[name] = duration
			sizes[name] = size
			fmt.Println()
		} else {
			fmt.Printf("Unknown converter: %s\n\n", name)
		}
	}

	fmt.Println("=== Summary ===")
	for _, name := range selectedConverters {
		if t, ok := results[name]; ok {
			size := sizes[name]
			sizeStr := fmt.Sprintf("%.1f KB", float64(size)/1024)
			fmt.Printf("%-10s %10v  %s\n", name+":", t, sizeStr)
		}
	}
}
