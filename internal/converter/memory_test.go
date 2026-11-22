package converter

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestProcessTreeRSS(t *testing.T) {
	proc := t.TempDir()
	files := map[string]string{
		"1/statm":            "20 10 0",
		"1/task/1/children":  "2",
		"1/task/11/children": "3",
		"2/statm":            "30 5 0",
		"2/task/2/children":  "4",
		"3/statm":            "40 7 0",
		"3/task/3/children":  "2",
		"4/statm":            "50 3 0",
		"4/task/4/children":  "",
	}
	for name, contents := range files {
		path := filepath.Join(proc, name)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if got := processTreeRSS(proc, 1, 4096); got != 25*4096 {
		t.Fatalf("tree RSS = %d, want %d", got, 25*4096)
	}
}

func TestRunProcessSamplesMemory(t *testing.T) {
	_, _, memory, err := runProcess([]string{"/bin/sh", "-c", "sleep 0.2"}, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if memory <= 0 {
		t.Fatal("expected a nonzero sampled RSS")
	}
}

func TestFormatPeakRSS(t *testing.T) {
	if FormatPeakRSS(0) != "FAILED" || FormatPeakRSS(12345678) != "12.35MB" {
		t.Fatal("incorrect peak RSS formatting")
	}
}
