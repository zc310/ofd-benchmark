package benchmarkdata

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"go.yaml.in/yaml/v4"
)

func TestFormatSize(t *testing.T) {
	for _, tt := range []struct {
		name string
		size int64
		pdf  string
		ofd  string
	}{
		{"zero", 0, "0.00KB", "0.00K"},
		{"whole KB", 1024, "1.00KB", "1.00K"},
		{"fractional KB", 1264, "1.23KB", "1.23K"},
		{"round up KB", 1275, "1.25KB", "1.25K"},
		{"below MB", 1024*1024 - 1, "1024.00KB", "1024.00K"},
		{"whole MB", 1024 * 1024, "1.00MB", "1.00M"},
		{"fractional MB", 1300000, "1.24MB", "1.24M"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatPDFSize(tt.size); got != tt.pdf {
				t.Errorf("FormatPDFSize(%d) = %q, want %q", tt.size, got, tt.pdf)
			}
			if got := FormatOFDSize(tt.size); got != tt.ofd {
				t.Errorf("FormatOFDSize(%d) = %q, want %q", tt.size, got, tt.ofd)
			}
		})
	}
}

func TestUpdateCell(t *testing.T) {
	path := copyBenchmarkYAML(t)
	want := readResultValues(t, path, "pdf-to-ofd", "pdf-speed", "go-zc310")
	want[0] = "42ms"
	updates := []CellUpdate{{
		Report: "pdf-to-ofd", Table: "pdf-speed", Input: "hello.pdf",
		Converter: "go-zc310", Value: "42ms",
	}}
	if err := Update(path, updates, false); err != nil {
		t.Fatal(err)
	}
	got := readResultValues(t, path, "pdf-to-ofd", "pdf-speed", "go-zc310")
	if !slices.Equal(got, want) {
		t.Fatalf("updated values = %v, want %v", got, want)
	}
}

func TestUpdateColumnWithReferenceCountInHeader(t *testing.T) {
	path := copyBenchmarkYAML(t)
	want := readResultValues(t, path, "pdf-to-ofd", "pdf-text-extraction", "go-zc310")
	want[0] = "31"
	updates := []CellUpdate{{
		Report: "pdf-to-ofd", Table: "pdf-text-extraction", Input: "hello.pdf",
		Converter: "go-zc310", Value: "31",
	}}
	if err := Update(path, updates, false); err != nil {
		t.Fatal(err)
	}
	got := readResultValues(t, path, "pdf-to-ofd", "pdf-text-extraction", "go-zc310")
	if !slices.Equal(got, want) {
		t.Fatalf("updated values for header with reference count = %v, want %v", got, want)
	}
}

func readResultValues(t *testing.T, path, reportID, tableID, label string) []string {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Reports []struct {
			ID     string `yaml:"id"`
			Tables []struct {
				ID   string `yaml:"id"`
				Rows []struct {
					Label  string   `yaml:"label"`
					Values []string `yaml:"values"`
				} `yaml:"rows"`
			} `yaml:"tables"`
		} `yaml:"reports"`
	}
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	for _, report := range doc.Reports {
		if report.ID != reportID {
			continue
		}
		for _, table := range report.Tables {
			if table.ID != tableID {
				continue
			}
			for _, row := range table.Rows {
				if row.Label == label && len(row.Values) > 0 {
					return row.Values
				}
			}
		}
	}
	t.Fatalf("missing result row %s/%s/%s", reportID, tableID, label)
	return nil
}

func TestResetKeepsReferenceBaseline(t *testing.T) {
	path := copyBenchmarkYAML(t)
	if err := Update(path, nil, true); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	text := string(raw)
	if !strings.Contains(text, "label: 数科版式阅读器, values: [1.00, 1.00, 1.00") {
		t.Fatal("reference baseline was reset")
	}
	if !strings.Contains(text, "values: [FAILED, FAILED, FAILED") {
		t.Fatal("benchmark cells were not reset")
	}
}

func TestWriteAndReadVersions(t *testing.T) {
	path := copyBenchmarkYAML(t)
	if err := WriteVersions(path, "v1-zc", "v1-ofdgo", "2.6.0"); err != nil {
		t.Fatal(err)
	}
	versions, err := ReadVersions(path)
	if err != nil {
		t.Fatal(err)
	}
	if versions[ZC310Module] != "v1-zc" || versions[OfdgoModule] != "v1-ofdgo" || versions[OfdrwArtifact] != "2.6.0" {
		t.Fatalf("unexpected versions: %#v", versions)
	}
}

func TestWriteEnvironment(t *testing.T) {
	path := copyBenchmarkYAML(t)
	values := map[string]string{
		EnvironmentOS:     "Ubuntu test",
		EnvironmentCPU:    "CPU test",
		EnvironmentCores:  "8",
		EnvironmentMemory: "16 GiB",
	}
	if err := WriteEnvironment(path, values); err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Environment map[string]string `yaml:"benchmark_environment"`
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	for key, want := range values {
		if doc.Environment[key] != want {
			t.Fatalf("environment %s = %q, want %q", key, doc.Environment[key], want)
		}
	}
}

func copyBenchmarkYAML(t *testing.T) string {
	t.Helper()
	input, err := os.ReadFile("../../docs/benchmark-results.yaml")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "benchmark-results.yaml")
	if err := os.WriteFile(path, input, 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}
