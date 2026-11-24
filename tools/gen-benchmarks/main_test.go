package main

import "testing"

func TestRenderTablePreservesData(t *testing.T) {
	table := Table{
		ID:      "measurements",
		Headers: []string{"转换器", "耗时", "相似度", "字符数"},
		Winner:  "none",
		Rows: []Row{
			{Label: "second", Values: []string{"1.20s", "99.00%", "0"}},
			{Label: "first", Values: []string{"23ms", "100.00%", "22"}},
			{Label: "failed", Values: []string{"FAILED", "FAILED", "FAILED"}},
		},
	}
	want := "| 转换器 | 耗时 | 相似度 | 字符数 |\n" +
		"|---|---:|---:|---:|\n" +
		"| second | 1.20s | 99.00% | 0 |\n" +
		"| first | 23ms | 100.00% | 22 |\n" +
		"| failed | FAILED | FAILED | FAILED |"
	got, err := renderTable(table)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("renderTable() = %q, want %q", got, want)
	}
}

func TestRenderTableKeepsDataBackedStatistics(t *testing.T) {
	table := Table{
		ID:      "speed",
		Headers: []string{"转换器", "a.ofd", "b.ofd"},
		Winner:  "min",
		Wins:    true,
		Rows: []Row{
			{Label: "first", Values: []string{"1.20s", "10ms"}},
			{Label: "second", Values: []string{"23ms", "20ms"}},
			{Label: "failed", Values: []string{"FAILED", "FAILED"}},
		},
	}
	want := "| 转换器 | a.ofd | b.ofd | 胜出次数 |\n" +
		"|---|---:|---:|---:|\n" +
		"| first | 1.20s | **10ms** | 1 |\n" +
		"| second | **23ms** | 20ms | 1 |\n" +
		"| failed | FAILED | FAILED | 0 |"
	got, err := renderTable(table)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("renderTable() = %q, want %q", got, want)
	}
}
