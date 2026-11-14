package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"ofd-benchmark/internal/benchmarkdata"
)

func main() {
	path := flag.String("yaml", "docs/benchmark-results.yaml", "benchmark results YAML")
	input := flag.String("input", "", "JSON cell updates file, or - for stdin")
	reset := flag.Bool("reset", false, "mark all measured cells FAILED before applying updates")
	flag.Parse()

	var updates []benchmarkdata.CellUpdate
	if *input != "" && *input != "-" {
		file, err := os.Open(*input)
		if err != nil {
			fatal(err)
		}
		defer file.Close()
		if err := decode(file, &updates); err != nil {
			fatal(err)
		}
	} else if *input == "-" {
		if err := decode(os.Stdin, &updates); err != nil {
			fatal(err)
		}
	}
	for _, update := range updates {
		if err := benchmarkdata.ValidateCell(update); err != nil {
			fatal(err)
		}
	}
	if err := benchmarkdata.Update(*path, updates, *reset); err != nil {
		fatal(err)
	}
}

func decode(r io.Reader, updates *[]benchmarkdata.CellUpdate) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(updates); err != nil {
		return err
	}
	return nil
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "update-benchmark-results:", err)
	os.Exit(1)
}
