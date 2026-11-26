package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"
	"go.yaml.in/yaml/v4"
)

type Source struct {
	BenchmarkVersions    map[string]string `yaml:"benchmark_versions"`
	BenchmarkEnvironment map[string]string `yaml:"benchmark_environment"`
	Reports              []Report          `yaml:"reports"`
}

type Report struct {
	ID     string  `yaml:"id"`
	Tables []Table `yaml:"tables"`
}

type Table struct {
	ID       string            `yaml:"id"`
	Headers  []string          `yaml:"headers"`
	Winner   string            `yaml:"winner"`
	Wins     bool              `yaml:"wins"`
	TieBreak map[string]string `yaml:"tie-break"`
	Rows     []Row             `yaml:"rows"`
}

type Row struct {
	Label  string   `yaml:"label"`
	Values []string `yaml:"values"`
}

var valuePattern = regexp.MustCompile(`^([0-9]+(?:\.[0-9]+)?)(.*)$`)

func main() {
	srcPath := flag.String("src", "docs/benchmark-results.yaml", "结果数据 YAML")
	check := flag.Bool("check", false, "仅检查文档是否与 YAML 一致")
	flag.Parse()

	raw, err := os.ReadFile(*srcPath)
	if err != nil {
		fatal(err)
	}
	var source Source
	if err := yaml.Unmarshal(raw, &source); err != nil {
		fatal(fmt.Errorf("%s: %w", *srcPath, err))
	}
	if err := validate(source); err != nil {
		fatal(fmt.Errorf("%s: %w", *srcPath, err))
	}

	for _, report := range source.Reports {
		path := "docs/" + report.ID + ".md"
		contents, err := os.ReadFile(path)
		if err != nil {
			fatal(err)
		}
		updated, err := updateDocument(string(contents), report, source.BenchmarkVersions, source.BenchmarkEnvironment)
		if err != nil {
			fatal(fmt.Errorf("%s: %w", path, err))
		}
		if *check {
			if updated != string(contents) {
				fatal(fmt.Errorf("%s: generated tables out of date (run go run ./tools/gen-benchmarks)", path))
			}
			continue
		}
		if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
			fatal(err)
		}
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "gen-benchmarks:", err)
	os.Exit(1)
}

func validate(source Source) error {
	if len(source.Reports) == 0 {
		return errors.New("reports 为空")
	}
	for _, key := range []string{"github.com/zc310/ofd", "github.com/xiaoqidun/ofdgo", "github.com/lihongjie0209/ofdrw-rs/bindings/go", "org.ofdrw:ofdrw-converter"} {
		if source.BenchmarkVersions[key] == "" {
			return fmt.Errorf("benchmark_versions 缺少 %s", key)
		}
	}
	for _, key := range []string{"os", "cpu_model", "logical_cores", "memory"} {
		if source.BenchmarkEnvironment[key] == "" {
			return fmt.Errorf("benchmark_environment 缺少 %s", key)
		}
	}
	seenReports := map[string]bool{}
	seenTables := map[string]bool{}
	for _, report := range source.Reports {
		if report.ID == "" || seenReports[report.ID] {
			return fmt.Errorf("report id 为空或重复: %q", report.ID)
		}
		seenReports[report.ID] = true
		for _, table := range report.Tables {
			if table.ID == "" || seenTables[table.ID] {
				return fmt.Errorf("table id 为空或重复: %q", table.ID)
			}
			seenTables[table.ID] = true
			if len(table.Headers) < 2 || len(table.Rows) == 0 {
				return fmt.Errorf("%s: 至少需要一个数据列和一行数据", table.ID)
			}
			if table.Winner != "min" && table.Winner != "max" && table.Winner != "none" {
				return fmt.Errorf("%s: winner 必须为 min/max/none", table.ID)
			}
			for _, row := range table.Rows {
				if row.Label == "" || len(row.Values) != len(table.Headers)-1 {
					return fmt.Errorf("%s: %q 的 values 数量与表头不符", table.ID, row.Label)
				}
				for _, value := range row.Values {
					if table.Winner != "none" && value != "FAILED" {
						if _, err := numericValue(value); err != nil {
							return fmt.Errorf("%s: 无法比较 %q: %w", table.ID, value, err)
						}
					}
				}
			}
		}
	}
	return nil
}

func updateDocument(doc string, report Report, versions, environment map[string]string) (string, error) {
	for _, table := range report.Tables {
		start := "<!-- benchmark-table:" + table.ID + ":start -->"
		end := "<!-- benchmark-table:" + table.ID + ":end -->"
		if strings.Count(doc, start) != 1 || strings.Count(doc, end) != 1 {
			return "", fmt.Errorf("table %s must have exactly one start/end marker", table.ID)
		}
		startAt := strings.Index(doc, start) + len(start)
		endAt := strings.Index(doc, end)
		if endAt < startAt {
			return "", fmt.Errorf("table %s markers are reversed", table.ID)
		}
		rendered, err := renderTable(table)
		if err != nil {
			return "", err
		}
		doc = doc[:startAt] + "\n\n" + rendered + "\n" + doc[endAt:]
	}
	chartSpecs := map[string]struct{ marker, table, title string }{
		"ofd-to-png": {"png", "png-speed", "OFD → PNG"},
		"ofd-to-pdf": {"ofd", "ofd-speed", "OFD → PDF"},
		"pdf-to-ofd": {"pdf", "pdf-speed", "PDF → OFD"},
	}
	if spec, ok := chartSpecs[report.ID]; ok {
		startMarker := "<!-- benchmark-charts:" + spec.marker + ":start -->"
		endMarker := "<!-- benchmark-charts:" + spec.marker + ":end -->"
		if strings.Count(doc, startMarker) != 1 || strings.Count(doc, endMarker) != 1 {
			return "", fmt.Errorf("%s chart markers missing or repeated", report.ID)
		}
		start := strings.Index(doc, startMarker) + len(startMarker)
		end := strings.Index(doc, endMarker)
		if end < start {
			return "", fmt.Errorf("%s chart markers are reversed", report.ID)
		}
		charts, err := renderSpeedChart(report.Tables, spec.table, spec.title)
		if err != nil {
			return "", err
		}
		doc = doc[:start] + "\n\n" + charts + "\n" + doc[end:]
	}
	if err := updateLibraryVersions(&doc, report.ID, versions); err != nil {
		return "", err
	}
	version := versions["org.ofdrw:ofdrw-converter"]
	if report.ID == "ofd-to-pdf" {
		if err := replaceVersionOnce(&doc, regexp.MustCompile("(org\\.ofdrw:ofdrw-converter:)[^`]+"), version); err != nil {
			return "", err
		}
	}
	startMarker := "<!-- benchmark-environment:start -->"
	endMarker := "<!-- benchmark-environment:end -->"
	if strings.Count(doc, startMarker) != 1 || strings.Count(doc, endMarker) != 1 {
		return "", fmt.Errorf("benchmark environment markers missing or repeated")
	}
	start := strings.Index(doc, startMarker) + len(startMarker)
	end := strings.Index(doc, endMarker)
	if end < start {
		return "", fmt.Errorf("benchmark environment markers are reversed")
	}
	block := fmt.Sprintf("\n\n- 系统：%s\n- CPU：%s（%s 逻辑核心）\n- 内存：%s\n", environment["os"], environment["cpu_model"], environment["logical_cores"], environment["memory"])
	doc = doc[:start] + block + doc[end:]
	return doc, nil
}

var markdownLink = regexp.MustCompile(`^\[([^\]]+)\]\([^)]+\)$`)

func renderSpeedChart(tables []Table, tableID, title string) (string, error) {
	var table *Table
	for i := range tables {
		if tables[i].ID == tableID {
			table = &tables[i]
			break
		}
	}
	if table == nil {
		return "", fmt.Errorf("missing chart source %s", tableID)
	}
	var labels, values, missing []string
	maximum := 0.0
	for _, row := range table.Rows {
		label := row.Label
		if match := markdownLink.FindStringSubmatch(label); match != nil {
			label = match[1]
		}
		total := 0.0
		var failed []string
		for col, sample := range table.Headers[1:] {
			value := row.Values[col]
			if value == "FAILED" {
				failed = append(failed, sample)
				continue
			}
			n, err := numericValue(value)
			if err != nil {
				return "", fmt.Errorf("%s %s %s: %w", tableID, sample, label, err)
			}
			total += n
		}
		if len(failed) > 0 {
			missing = append(missing, "`"+label+"`（"+strings.Join(failed, "、")+"）")
			continue
		}
		total /= 1000
		maximum = math.Max(maximum, total)
		labels = append(labels, strconv.Quote(label))
		values = append(values, strconv.FormatFloat(total, 'f', -1, 64))
	}
	var b strings.Builder
	if len(values) > 0 {
		fmt.Fprintf(&b, "```mermaid\n---\nconfig:\n  xyChart:\n    height: %d\n---\nxychart-beta horizontal\n    title %s\n    x-axis [%s]\n    y-axis \"秒\" 0 --> %s\n    bar [%s]\n```\n\n",
			max(500, len(values)*45+150), strconv.Quote(title+" 全部文件总耗时（秒，越低越快）"), strings.Join(labels, ", "),
			strconv.FormatFloat(math.Max(1, math.Ceil(maximum)), 'f', -1, 64), strings.Join(values, ", "))
	} else {
		b.WriteString("暂无有效数据。\n\n")
	}
	if len(missing) > 0 {
		fmt.Fprintf(&b, "未完成全部文件（括号内为无有效结果的文件）：%s。\n\n", strings.Join(missing, "；"))
	}
	return strings.TrimSpace(b.String()), nil
}

func updateLibraryVersions(doc *string, reportID string, versions map[string]string) error {
	zc310 := versions["github.com/zc310/ofd"]
	ofdgo := versions["github.com/xiaoqidun/ofdgo"]
	ofdrwRS := versions["github.com/lihongjie0209/ofdrw-rs/bindings/go"]
	ofdrw := versions["org.ofdrw:ofdrw-converter"]
	if reportID == "ofd-to-png" {
		for _, update := range []struct{ prefix, version string }{
			{"| [go-zc310]", zc310},
			{"| [go-ofdgo]", ofdgo},
			{"| [rust-ofd-utility]", versions["github.com/ofd-utility/ofd-utility"]},
			{"| [go-ofdrw-rs]", ofdrwRS},
			{"| [java-ofdrw]", ofdrw},
			{"| [rust-ofdmanager]", versions["github.com/feuvan/ofdmanager"]},
		} {
			if update.version == "" {
				return fmt.Errorf("missing PNG converter version for %s", update.prefix)
			}
			// Match only the version table, not generated result rows.
			if err := replaceMarkdownRowCell(doc, update.prefix, "200 DPI", 2, update.version); err != nil {
				return err
			}
		}
	}
	if reportID == "ofd-to-pdf" {
		for _, update := range []struct {
			prefix, marker string
			column         int
			version        string
		}{
			{"| [zc310-ofd]", "", 2, zc310},
			{"| [xiaoqidun-ofdgo]", "", 2, ofdgo},
			{"| [ofdrw-rs Go binding]", "✅", 2, ofdrwRS},
			{"| [ofdrw](https://github.com/ofdrw/ofdrw)", "✅", 2, ofdrw},
		} {
			if err := replaceMarkdownRowCell(doc, update.prefix, update.marker, update.column, update.version); err != nil {
				return err
			}
		}
	}
	if reportID == "pdf-to-ofd" {
		for _, update := range []struct {
			prefix, marker string
			column         int
			version        string
		}{
			{"| go-zc310", "zc310/ofd", 4, zc310},
			{"| go-ofdgo", "ConvertPDF", 4, ofdgo},
			{"| java-ofdrw", "PDFConverter", 4, ofdrw},
		} {
			if err := replaceMarkdownRowCell(doc, update.prefix, update.marker, update.column, update.version); err != nil {
				return err
			}
		}
	}
	return nil
}

func replaceMarkdownRowCell(doc *string, rowPrefix, marker string, cellIndex int, value string) error {
	lines := strings.SplitAfter(*doc, "\n")
	matches := 0
	for i, line := range lines {
		if !strings.HasPrefix(line, rowPrefix) || !strings.Contains(line, marker) {
			continue
		}
		cells := strings.Split(line, "|")
		if cellIndex >= len(cells)-1 {
			return fmt.Errorf("malformed Markdown row for %s", rowPrefix)
		}
		cells[cellIndex] = " " + value + " "
		lines[i] = strings.Join(cells, "|")
		matches++
	}
	if matches != 1 {
		return fmt.Errorf("expected one Markdown version row for %s, found %d", rowPrefix, matches)
	}
	*doc = strings.Join(lines, "")
	return nil
}

func replaceVersionOnce(text *string, pattern *regexp.Regexp, version string) error {
	matches := pattern.FindAllStringSubmatchIndex(*text, -1)
	if len(matches) != 1 {
		return fmt.Errorf("expected one version marker for %s, found %d", version, len(matches))
	}
	match := matches[0]
	*text = (*text)[:match[3]] + version + (*text)[match[1]:]
	return nil
}

func renderTable(table Table) (string, error) {
	twinning := make([]map[int]bool, len(table.Rows))
	for i := range twinning {
		twinning[i] = make(map[int]bool)
	}
	if table.Winner != "none" {
		for col := range table.Headers[1:] {
			best := math.NaN()
			for rowIndex, row := range table.Rows {
				value := row.Values[col]
				if value == "FAILED" {
					continue
				}
				n, _ := numericValue(value)
				if math.IsNaN(best) || table.Winner == "min" && n < best || table.Winner == "max" && n > best {
					best = n
					for i := range twinning {
						delete(twinning[i], col)
					}
					twinning[rowIndex][col] = true
				} else if n == best {
					twinning[rowIndex][col] = true
				}
			}
		}
	}

	// Keep ties bold in every row, but assign the win counter to one declared row.
	winCounts := make([]int, len(table.Rows))
	for col := range table.Headers[1:] {
		winners := []int{}
		for rowIndex := range table.Rows {
			if twinning[rowIndex][col] {
				winners = append(winners, rowIndex)
			}
		}
		if len(winners) == 0 {
			continue
		}
		winner := winners[0]
		if preferred, ok := table.TieBreak[table.Headers[col+1]]; ok && len(winners) > 1 {
			for _, rowIndex := range winners {
				if table.Rows[rowIndex].Label == preferred {
					winner = rowIndex
					break
				}
			}
		}
		winCounts[winner]++
	}

	hasPercentage := false
	var b strings.Builder
	b.WriteString("| " + strings.Join(table.Headers, " | "))
	if table.Wins {
		b.WriteString(" | 胜出次数")
	}
	b.WriteString(" |\n|---")
	for range table.Headers[1:] {
		b.WriteString("|---:")
	}
	if table.Wins {
		b.WriteString("|---:")
	}
	b.WriteString("|\n")
	for rowIndex, row := range table.Rows {
		cells := make([]string, len(row.Values)+1)
		cells[0] = row.Label
		for col, value := range row.Values {
			if strings.HasSuffix(value, "%") {
				hasPercentage = true
				n, err := numericValue(value)
				if err != nil {
					return "", err
				}
				value = fmt.Sprintf("%.4f%%", n)
			}
			if twinning[rowIndex][col] {
				value = "**" + value + "**"
			}
			cells[col+1] = value
		}
		b.WriteString("| " + strings.Join(cells, " | "))
		if table.Wins {
			b.WriteString(fmt.Sprintf(" | %d", winCounts[rowIndex]))
		}
		b.WriteString(" |\n")
	}
	rendered := strings.TrimSuffix(b.String(), "\n")
	if hasPercentage {
		rendered = alignMarkdownTable(rendered)
	}
	return rendered, nil
}

// Pad source cells too: converter names stay left-aligned and numeric cells
// right-aligned. Count display columns so Chinese headers line up with ASCII.
func alignMarkdownTable(table string) string {
	lines := strings.Split(table, "\n")
	rows := make([][]string, len(lines))
	var widths []int
	for i, line := range lines {
		rows[i] = strings.Split(strings.Trim(line, "|"), "|")
		if i == 0 {
			widths = make([]int, len(rows[i]))
		}
		for col, cell := range rows[i] {
			rows[i][col] = strings.TrimSpace(cell)
			if i != 1 {
				widths[col] = max(widths[col], runewidth.StringWidth(rows[i][col]), 4)
			}
		}
	}
	for i, row := range rows {
		for col, cell := range row {
			if i == 1 {
				cell = strings.Repeat("-", widths[col])
				if col > 0 {
					cell = cell[:len(cell)-1] + ":"
				}
			} else {
				padding := strings.Repeat(" ", widths[col]-runewidth.StringWidth(cell))
				if col == 0 {
					cell += padding
				} else {
					cell = padding + cell
				}
			}
			row[col] = cell
		}
		lines[i] = "| " + strings.Join(row, " | ") + " |"
	}
	return strings.Join(lines, "\n")
}

func numericValue(value string) (float64, error) {
	parts := valuePattern.FindStringSubmatch(value)
	if parts == nil {
		return 0, fmt.Errorf("invalid numeric value")
	}
	var number float64
	if _, err := fmt.Sscan(parts[1], &number); err != nil {
		return 0, err
	}
	switch strings.ToLower(parts[2]) {
	case "", "%":
		return number, nil
	case "ms":
		return number, nil
	case "s":
		return number * 1000, nil
	case "k", "kb":
		return number * 1000, nil
	case "m", "mb":
		return number * 1000 * 1000, nil
	default:
		return 0, fmt.Errorf("unknown unit %q", parts[2])
	}
}
