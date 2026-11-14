package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"ofd-benchmark/internal/benchmarkdata"
	"ofd-benchmark/internal/converter"
)

type Converter struct {
	Name       string
	Command    []string
	OutputFile string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ofd-benchmark <file.ofd> [converter1] [converter2] ...")
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
	outputDir := filepath.Join(binDir, "output")
	os.MkdirAll(outputDir, 0755)

	// Get input filename without extension
	inputBase := filepath.Base(inputFile)
	inputName := inputBase[:len(inputBase)-len(filepath.Ext(inputBase))]

	converters := map[string]Converter{
		"ofdgo": {
			Name:       "go-ofdgo",
			Command:    []string{filepath.Join(binDir, "ofdgo-convert"), inputFile, filepath.Join(binDir, "output_ofdgo.pdf")},
			OutputFile: filepath.Join(binDir, "output_ofdgo.pdf"),
		},
		"zc310": {
			Name:       "go-zc310",
			Command:    []string{filepath.Join(binDir, "zc310-convert"), inputFile, filepath.Join(binDir, "output_zc310.pdf")},
			OutputFile: filepath.Join(binDir, "output_zc310.pdf"),
		},
		"python": {
			Name:       "python-easyofd",
			Command:    []string{"python3", filepath.Join(binDir, "..", "python", "ofd2pdf.py"), inputFile, filepath.Join(binDir, "output_python.pdf")},
			OutputFile: filepath.Join(binDir, "output_python.pdf"),
		},
		"ofd2pdf": {
			Name:       "python-ofd2pdf",
			Command:    []string{"python3", filepath.Join(binDir, "..", "python", "ofd2pdf_ofd2pdf.py"), inputFile, filepath.Join(binDir, "output_ofd2pdf.pdf")},
			OutputFile: filepath.Join(binDir, "output_ofd2pdf.pdf"),
		},
		"java": {
			Name:       "java-ofdrw",
			Command:    []string{filepath.Join(binDir, "..", "java", "ofd2pdf_java.sh"), inputFile, filepath.Join(binDir, "output_java.pdf")},
			OutputFile: filepath.Join(binDir, "output_java.pdf"),
		},
		"node": {
			Name:       "node-ofd2pdf",
			Command:    []string{filepath.Join(binDir, "..", "node", "ofd2pdf_node.sh"), inputFile, filepath.Join(binDir, "output_node.pdf")},
			OutputFile: filepath.Join(binDir, "output_node.pdf"),
		},
	}

	selectedConverters := os.Args[2:]
	if len(selectedConverters) == 0 {
		selectedConverters = []string{"ofdgo", "zc310", "python", "ofd2pdf", "java", "node"}
	}

	fmt.Printf("Input file: %s\n\n", inputFile)

	type result struct {
		duration time.Duration
		size     int64
		ok       bool
	}
	results := make(map[string]result)

	for _, name := range selectedConverters {
		if c, ok := converters[name]; ok {
			duration, size, ok := converter.Run(c.Name, c.Command, c.OutputFile, timeout)
			results[c.Name] = result{duration, size, ok}

			if ok {
				// Move to output directory
				dstFile := filepath.Join(outputDir, fmt.Sprintf("%s_%s.pdf", inputName, c.Name))
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
					sizeStr := fmt.Sprintf("%.1f KB", float64(r.size)/1024)
					fmt.Printf("%-20s %10v  %s\n", c.Name+":", r.duration, sizeStr)
				} else {
					fmt.Printf("%-20s %10v  FAILED\n", c.Name+":", r.duration)
				}
			}
		}
	}

	projectDir := filepath.Dir(binDir)
	refPDF := filepath.Join(projectDir, "testdata", "suwell", inputName+".pdf")
	var refSize int64
	if info, err := os.Stat(refPDF); err == nil {
		refSize = info.Size()
	}
	var updates []benchmarkdata.CellUpdate
	for _, name := range selectedConverters {
		c, exists := converters[name]
		if !exists {
			continue
		}
		r, exists := results[c.Name]
		if !exists {
			continue
		}
		rowName, hasRow := map[string]string{
			"ofdgo": "go-ofdgo", "zc310": "go-zc310", "python": "easyofd",
			"ofd2pdf": "ofd2pdf",
			"java":    "ofdrw", "node": "@miconvert/ofd-to-pdf",
		}[name]
		if !hasRow {
			continue
		}
		elapsed, size, ratio := "FAILED", "FAILED", "FAILED"
		if r.ok {
			elapsed = benchmarkdata.FormatDuration(r.duration)
			size = benchmarkdata.FormatPDFSize(r.size)
			ratio = benchmarkdata.FormatRatio(r.size, refSize)
		}
		tables := map[string]string{
			"ofd-pdf-size":   size,
			"ofd-speed":      elapsed,
			"ofd-size-ratio": ratio,
		}
		for table, value := range tables {
			updates = append(updates, benchmarkdata.CellUpdate{
				Report: "ofd-to-pdf", Table: table, Input: inputBase,
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
