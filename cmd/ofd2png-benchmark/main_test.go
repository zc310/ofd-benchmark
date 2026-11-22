package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go.yaml.in/yaml/v4"
)

func writePNG(t *testing.T, path string) {
	t.Helper()
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	if err := png.Encode(file, image.NewRGBA(image.Rect(0, 0, 10, 10))); err != nil {
		t.Fatal(err)
	}
}

func TestRegistryMatchesReportGroups(t *testing.T) {
	raw, err := os.ReadFile("../../docs/benchmark-results.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var source struct {
		Reports []struct {
			ID     string
			Tables []struct {
				ID   string
				Rows []struct{ Label string }
			}
		}
	}
	if err := yaml.Unmarshal(raw, &source); err != nil {
		t.Fatal(err)
	}
	registry := pngConverters("/bin", "/project")
	native, rows := 0, 0
	for _, c := range registry {
		if c.tables == "png" {
			native++
		} else {
			t.Fatalf("unknown PNG group %q", c.tables)
		}
	}
	for _, report := range source.Reports {
		if report.ID != "ofd-to-png" {
			continue
		}
		for _, table := range report.Tables {
			for _, row := range table.Rows {
				matched := false
				for _, c := range registry {
					if strings.HasPrefix(row.Label, "["+c.name+"](") {
						matched = true
						if !strings.HasPrefix(table.ID, c.tables+"-") ||
							(c.tables == "png" && strings.HasPrefix(table.ID, "png-via-pdf-")) {
							t.Fatalf("%s in wrong table %s", c.name, table.ID)
						}
					}
				}
				if !matched {
					t.Fatalf("unregistered PNG report row %s", row.Label)
				}
				rows++
			}
		}
	}
	if native != 8 || rows != len(registry)*8 {
		t.Fatalf("native=%d report rows=%d", native, rows)
	}
	for _, excluded := range []string{"ofd2img", "python", "node", "easyofd-rust"} {
		if _, exists := registry[excluded]; exists {
			t.Fatalf("excluded PNG converter %s must not participate", excluded)
		}
	}
}

func TestInspectPages(t *testing.T) {
	dir := t.TempDir()
	if _, _, err := inspectPages(dir, "page"); err == nil {
		t.Fatal("empty output accepted")
	}
	for i := 1; i <= 12; i++ {
		writePNG(t, filepath.Join(dir, fmt.Sprintf("page_%d.png", i)))
	}
	pages, size, err := inspectPages(dir, "page")
	if err != nil || pages != 12 || size <= 0 {
		t.Fatalf("pages=%d size=%d err=%v", pages, size, err)
	}
	if err := os.Remove(filepath.Join(dir, "page_2.png")); err != nil {
		t.Fatal(err)
	}
	if _, _, err := inspectPages(dir, "page"); err == nil {
		t.Fatal("missing page accepted")
	}
}

func TestInspectPagesRejectsCorruptionAndUnexpectedFiles(t *testing.T) {
	for _, name := range []string{"page_1.png", "other.png", "page_0.png", "page_01.png"} {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			if err := os.WriteFile(filepath.Join(dir, name), []byte("not PNG"), 0o644); err != nil {
				t.Fatal(err)
			}
			if _, _, err := inspectPages(dir, "page"); err == nil {
				t.Fatal("invalid output accepted")
			}
		})
	}
}
