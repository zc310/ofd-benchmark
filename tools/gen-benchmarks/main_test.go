package main

import (
	"strings"
	"testing"
)

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
	want := "| 转换器 |   耗时 |    相似度 | 字符数 |\n" +
		"| ------ | -----: | --------: | -----: |\n" +
		"| second |  1.20s |  99.0000% |      0 |\n" +
		"| first  |   23ms | 100.0000% |     22 |\n" +
		"| failed | FAILED |    FAILED | FAILED |"
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

func TestRenderTableFourDecimalPercentages(t *testing.T) {
	table := Table{
		ID:      "precision",
		Headers: []string{"转换器", "a.ofd", "b.ofd"},
		Winner:  "max",
		Wins:    true,
		Rows: []Row{
			{Label: "first", Values: []string{"99.9960%", "98.12%"}},
			{Label: "second", Values: []string{"99.9961%", "FAILED"}},
		},
	}
	want := "| 转换器 |        a.ofd |        b.ofd | 胜出次数 |\n" +
		"| ------ | -----------: | -----------: | -------: |\n" +
		"| first  |     99.9960% | **98.1200%** |        1 |\n" +
		"| second | **99.9961%** |       FAILED |        1 |"
	got, err := renderTable(table)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("renderTable() = %q, want %q", got, want)
	}
	if table.Rows[0].Values[1] != "98.12%" {
		t.Fatal("renderTable changed historical source precision")
	}
}

func TestAlignMarkdownTableIsIdempotent(t *testing.T) {
	table := "| 转换器 | 比例 |\n|---|---:|\n| first | **99.9961%** |\n| second | FAILED |"
	aligned := alignMarkdownTable(table)
	if got := alignMarkdownTable(aligned); got != aligned {
		t.Fatalf("alignment changed on second pass:\n%s", got)
	}
}

func TestPNGChartsConvertUnitsAndOmitFailures(t *testing.T) {
	tables := []Table{
		{ID: "png-speed", Headers: []string{"转换器", "a.ofd", "b.ofd"}, Rows: []Row{
			{Label: "[first](https://example.com)", Values: []string{"523ms", "1.20s"}},
			{Label: "second", Values: []string{"1.20s", "300ms"}},
			{Label: "failed", Values: []string{"10ms", "FAILED"}},
		}},
	}
	got, err := renderSpeedChart(tables, "png-speed", "OFD → PNG")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`xychart-beta horizontal`,
		`x-axis ["first", "second"]`,
		`bar [1.723, 1.5]`,
		"未完成全部文件（括号内为无有效结果的文件）：`failed`（b.ofd）。",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("chart missing %q:\n%s", want, got)
		}
	}
}
