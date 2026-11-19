package main

import (
	"context"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/zc310/ofd/pkg/converter"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: zc310-png <input.ofd> <output-directory> (200 DPI)")
	}
	if err := os.MkdirAll(os.Args[2], 0o755); err != nil {
		log.Fatal(err)
	}
	err := converter.Image(context.Background(), os.Args[1],
		converter.PNG(), converter.DPI(200), converter.BgColor(color.White),
		converter.Writer(func(page int) (io.WriteCloser, error) {
			return os.Create(filepath.Join(os.Args[2], fmt.Sprintf("page_%d.png", page)))
		}))
	if err != nil {
		log.Fatal(err)
	}
}
