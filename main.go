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

func runConverter(name string, args []string, outputFile string) (time.Duration, int64, bool) {
	fmt.Printf("=== %s ===\n", name)
	start := time.Now()

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = os.TempDir()

	err := cmd.Run()
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("FAILED: %v\n", err)
		return duration, 0, false
	}

	var size int64
	if info, err := os.Stat(outputFile); err == nil {
		size = info.Size()
	} else {
		fmt.Printf("FAILED: output file not created\n")
		return duration, 0, false
	}

	fmt.Printf("Time: %v\n", duration)
	return duration, size, true
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
		"rust": {
			Name:       "rust-easyofd",
			Command:    []string{filepath.Join(binDir, "easyofd"), inputFile, filepath.Join(binDir, "output_rust.pdf")},
			OutputFile: filepath.Join(binDir, "output_rust.pdf"),
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
		"ofdreader-python": {
			Name:       "python-ofdreader",
			Command:    []string{"python3", filepath.Join(binDir, "..", "python", "ofd2pdf_ofdreader.py"), inputFile, filepath.Join(binDir, "output_ofdreader-python.pdf")},
			OutputFile: filepath.Join(binDir, "output_ofdreader-python.pdf"),
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
		selectedConverters = []string{"ofdgo", "zc310", "rust", "python", "ofd2pdf", "ofdreader-python", "java", "node"}
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
			duration, size, ok := runConverter(c.Name, c.Command, c.OutputFile)
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
}
