package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/xiaoqidun/ofdgo"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ofdgo-png <input.ofd> <output-directory> (200 DPI)")
	}
	if err := convert(os.Args[1], os.Args[2]); err != nil {
		log.Fatal(err)
	}
}

func convert(inputFile, outputDir string) error {
	input, err := ofdgo.OpenFile(context.Background(), inputFile, ofdgo.InputOptions{})
	if err != nil {
		return err
	}
	defer input.Close()
	reader, err := ofdgo.NewReader(input)
	if err != nil {
		return err
	}
	defer reader.Close()
	count, err := reader.PageCount()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}
	renderer := ofdgo.NewRenderer(reader, ofdgo.WithDPI(200), ofdgo.WithAnnotations(true))
	for i := 0; i < count; i++ {
		page, err := reader.PageContentByIndex(i)
		if err != nil {
			return err
		}
		file, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("page_%d.png", i+1)))
		if err != nil {
			return err
		}
		renderErr := renderer.RenderToPNG(page, file)
		closeErr := file.Close()
		if renderErr != nil {
			return renderErr
		}
		if closeErr != nil {
			return closeErr
		}
	}
	return nil
}
