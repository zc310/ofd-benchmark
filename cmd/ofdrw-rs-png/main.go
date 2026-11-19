package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/lihongjie0209/ofdrw-rs/bindings/go/ofdrw"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ofdrw-rs-png <input.ofd> <output-directory> (200 DPI)")
	}
	if os.Getenv("OFDRW_LIB_PATH") == "" {
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		library := map[string]string{"linux": "libofdrw_ffi.so", "darwin": "libofdrw_ffi.dylib", "windows": "ofdrw_ffi.dll"}[runtime.GOOS]
		if library == "" {
			log.Fatalf("unsupported platform: %s", runtime.GOOS)
		}
		if err := os.Setenv("OFDRW_LIB_PATH", filepath.Join(filepath.Dir(exe), "ofdrw-rs", library)); err != nil {
			log.Fatal(err)
		}
	}
	if err := convert(os.Args[1], os.Args[2]); err != nil {
		log.Fatal(err)
	}
}

func convert(input, output string) error {
	doc, err := ofdrw.Open(input)
	if err != nil {
		return err
	}
	defer doc.Close()
	count, err := doc.PageCount()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(output, 0o755); err != nil {
		return err
	}
	for i := 0; i < count; i++ {
		data, err := doc.PagePNG(i, 200)
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(output, fmt.Sprintf("page_%d.png", i+1)), data, 0o644); err != nil {
			return err
		}
	}
	return nil
}
