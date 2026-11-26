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

func TestPNGSpeedSVG(t *testing.T) {
	tables := []Table{{ID: "png-speed", Headers: []string{"库", "a", "b"}, Rows: []Row{
		{Label: "slow", Values: []string{"1s", "500ms"}},
		{Label: "fast & safe", Values: []string{"100ms", "200ms"}},
		{Label: "incomplete", Values: []string{"1ms", "FAILED"}},
	}}}
	svg, missing, err := renderChartSVG(tables, chartSpec{table: "png-speed", title: "OFD → PNG"}, "speed")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(svg, "fast &amp; safe") || !strings.Contains(svg, "0.300 s") ||
		strings.Index(svg, "fast &amp; safe") > strings.Index(svg, ">slow</text>") ||
		strings.Contains(svg, ">incomplete</text>") || len(missing) != 1 || missing[0] != "`incomplete`（b）" {
		t.Fatalf("unexpected chart or missing results: %s, %v", svg, missing)
	}
}

func TestSizeChartTotalsAndMissing(t *testing.T) {
	tables := []Table{{ID: "size", Headers: []string{"库", "a", "b"}, Rows: []Row{
		{Label: "large", Values: []string{"2MB", "500KB"}},
		{Label: "small", Values: []string{"200KB", "100KB"}},
		{Label: "incomplete", Values: []string{"1MB", "FAILED"}},
	}}}
	svg, missing, err := renderChartSVG(tables, chartSpec{sizeTable: "size", title: "test"}, "size")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(svg, "0.300 MB") || !strings.Contains(svg, "2.500 MB") ||
		strings.Index(svg, ">small</text>") > strings.Index(svg, ">large</text>") ||
		strings.Contains(svg, ">incomplete</text>") || len(missing) != 1 || missing[0] != "`incomplete`（b）" {
		t.Fatalf("unexpected size chart or missing results: %s, %v", svg, missing)
	}
}

func TestCombinedChartIndependentMissingValues(t *testing.T) {
	tables := []Table{
		{ID: "speed", Headers: []string{"库", "a", "b"}, Rows: []Row{
			{Label: "first", Values: []string{"1s", "500ms"}},
			{Label: "second", Values: []string{"FAILED", "1s"}},
		}},
		{ID: "size", Headers: []string{"库", "a", "b"}, Rows: []Row{
			{Label: "first", Values: []string{"FAILED", "1MB"}},
			{Label: "second", Values: []string{"500KB", "1MB"}},
		}},
	}
	svg, err := renderCombinedSVG(tables, chartSpec{table: "speed", sizeTable: "size"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(svg, "无完整结果") != 2 || strings.Count(svg, ">1.500</text>") != 2 {
		t.Fatalf("unexpected combined chart: %s", svg)
	}
}
