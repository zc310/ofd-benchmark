package benchmarkdata

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"go.yaml.in/yaml/v4"
)

type CellUpdate struct {
	Report    string `json:"report"`
	Table     string `json:"table"`
	Input     string `json:"input"`
	Converter string `json:"converter"`
	Value     string `json:"value"`
}

func FormatDuration(duration time.Duration) string {
	if duration < time.Second {
		return fmt.Sprintf("%dms", duration.Round(time.Millisecond)/time.Millisecond)
	}
	return fmt.Sprintf("%.2fs", float64(duration)/float64(time.Second))
}

func FormatPDFSize(size int64) string {
	kb := float64(size) / 1024
	if kb >= 1024 {
		return fmt.Sprintf("%.2fMB", kb/1024)
	}
	return fmt.Sprintf("%.2fKB", kb)
}

func FormatOFDSize(size int64) string {
	kb := float64(size) / 1024
	if kb >= 1024 {
		return fmt.Sprintf("%.2fM", kb/1024)
	}
	return fmt.Sprintf("%.2fK", kb)
}

func FormatRatio(size, reference int64) string {
	if reference <= 0 {
		return "FAILED"
	}
	return fmt.Sprintf("%.2f", float64(size)/float64(reference))
}

const (
	ZC310Module       = "github.com/zc310/ofd"
	OfdgoModule       = "github.com/xiaoqidun/ofdgo"
	OfdrwArtifact     = "org.ofdrw:ofdrw-converter"
	EnvironmentOS     = "os"
	EnvironmentCPU    = "cpu_model"
	EnvironmentCores  = "logical_cores"
	EnvironmentMemory = "memory"
)

// Update applies result cells to the benchmark YAML under an exclusive lock,
// then atomically replaces the file so concurrent benchmark processes cannot
// truncate or interleave updates.
func Update(path string, updates []CellUpdate, reset bool) error {
	return mutate(path, func(root *yaml.Node) error {
		if reset {
			resetResults(root)
		}
		for _, update := range updates {
			if err := setCell(root, update); err != nil {
				return err
			}
		}
		return nil
	})
}

func ReadVersions(path string) (map[string]string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var doc yaml.Node
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	versions := make(map[string]string)
	versionsNode := mapValue(doc.Content[0], "benchmark_versions")
	if versionsNode == nil {
		return versions, nil
	}
	for i := 0; i+1 < len(versionsNode.Content); i += 2 {
		versions[versionsNode.Content[i].Value] = scalar(versionsNode.Content[i+1])
	}
	return versions, nil
}

func WriteVersions(path, zc310, ofdgo, ofdrw string) error {
	return mutate(path, func(root *yaml.Node) error {
		versions := mapValue(root, "benchmark_versions")
		if versions == nil {
			return fmt.Errorf("benchmark_versions missing from YAML")
		}
		setMappingScalar(versions, ZC310Module, zc310)
		setMappingScalar(versions, OfdgoModule, ofdgo)
		setMappingScalar(versions, OfdrwArtifact, ofdrw)
		return nil
	})
}

func WriteEnvironment(path string, values map[string]string) error {
	return mutate(path, func(root *yaml.Node) error {
		environment := mapValue(root, "benchmark_environment")
		if environment == nil {
			return fmt.Errorf("benchmark_environment missing from YAML")
		}
		for key, value := range values {
			setMappingScalar(environment, key, value)
		}
		return nil
	})
}

func mutate(path string, change func(root *yaml.Node) error) error {
	lock, err := os.OpenFile(path+".lock", os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return err
	}
	defer lock.Close()
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX); err != nil {
		return err
	}
	defer syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)

	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var doc yaml.Node
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		return fmt.Errorf("parse %s: %w", path, err)
	}
	if len(doc.Content) == 0 {
		return fmt.Errorf("%s: empty YAML document", path)
	}
	root := doc.Content[0]
	if err := change(root); err != nil {
		return err
	}

	encoded, err := yaml.Marshal(&doc)
	if err != nil {
		return err
	}
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".benchmark-results-*.yaml")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if err := tmp.Chmod(info.Mode().Perm()); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(encoded); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpName, path)
}

func setMappingScalar(mapping *yaml.Node, key, value string) {
	for i := 0; i+1 < len(mapping.Content); i += 2 {
		if mapping.Content[i].Value == key {
			setScalar(mapping.Content[i+1], value)
			return
		}
	}
	keyNode := &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: key}
	valueNode := &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: value}
	mapping.Content = append(mapping.Content, keyNode, valueNode)
}

func resetResults(root *yaml.Node) {
	for _, report := range sequence(mapValue(root, "reports")) {
		for _, table := range sequence(mapValue(report, "tables")) {
			for _, row := range sequence(mapValue(table, "rows")) {
				label := scalar(mapValue(row, "label"))
				if label == "数科版式阅读器" {
					continue
				}
				for _, value := range sequence(mapValue(row, "values")) {
					setScalar(value, "FAILED")
				}
			}
		}
	}
}

func setCell(root *yaml.Node, update CellUpdate) error {
	report := findByScalar(sequence(mapValue(root, "reports")), "id", update.Report)
	if report == nil {
		return fmt.Errorf("unknown report %q", update.Report)
	}
	table := findByScalar(sequence(mapValue(report, "tables")), "id", update.Table)
	if table == nil {
		return fmt.Errorf("report %s has no table %q", update.Report, update.Table)
	}
	headers := sequence(mapValue(table, "headers"))
	column := -1
	for i, header := range headers[1:] {
		if headerInput(scalar(header)) == update.Input {
			column = i
			break
		}
	}
	if column < 0 {
		return fmt.Errorf("table %s has no input column %q", update.Table, update.Input)
	}
	for _, row := range sequence(mapValue(table, "rows")) {
		label := scalar(mapValue(row, "label"))
		if converterKey(label) != update.Converter {
			continue
		}
		values := sequence(mapValue(row, "values"))
		if column >= len(values) {
			return fmt.Errorf("row %q in table %s has no value at column %d", label, update.Table, column)
		}
		setScalar(values[column], update.Value)
		return nil
	}
	return fmt.Errorf("table %s has no converter %q", update.Table, update.Converter)
}

func headerInput(header string) string {
	if i := strings.IndexByte(header, '('); i >= 0 {
		return header[:i]
	}
	return header
}

func converterKey(label string) string {
	if strings.HasPrefix(label, "[") {
		if end := strings.Index(label, "]("); end > 1 {
			return label[1:end]
		}
	}
	return label
}

func findByScalar(nodes []*yaml.Node, key, value string) *yaml.Node {
	for _, node := range nodes {
		if scalar(mapValue(node, key)) == value {
			return node
		}
	}
	return nil
}

func mapValue(node *yaml.Node, key string) *yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			return node.Content[i+1]
		}
	}
	return nil
}

func sequence(node *yaml.Node) []*yaml.Node {
	if node == nil || node.Kind != yaml.SequenceNode {
		return nil
	}
	return node.Content
}

func scalar(node *yaml.Node) string {
	if node == nil || node.Kind != yaml.ScalarNode {
		return ""
	}
	return node.Value
}

func setScalar(node *yaml.Node, value string) {
	if node == nil {
		return
	}
	node.Kind = yaml.ScalarNode
	node.Tag = "!!str"
	node.Value = value
	node.Style = 0
}

// ValidateCell provides a fast fail for incomplete caller data.
func ValidateCell(update CellUpdate) error {
	if update.Report == "" || update.Table == "" || update.Input == "" || update.Converter == "" || update.Value == "" {
		return errors.New("cell update is missing a required field")
	}
	return nil
}
