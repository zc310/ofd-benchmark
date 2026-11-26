package main

import (
	"errors"
	"flag"
	"fmt"
	"html"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
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

type chartSpec struct{ marker, table, sizeTable, title string }

var chartSpecs = map[string]chartSpec{
	"ofd-to-png": {"png", "png-speed", "png-size", "OFD → PNG"},
	"ofd-to-pdf": {"ofd", "ofd-speed", "ofd-pdf-size", "OFD → PDF"},
	"pdf-to-ofd": {"pdf", "pdf-speed", "pdf-ofd-size", "PDF → OFD"},
}

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
		if spec, ok := chartSpecs[report.ID]; ok {
			{
				svg, err := renderCombinedSVG(report.Tables, spec)
				if err != nil {
					fatal(err)
				}
				svgPath := "docs/charts/" + report.ID + "-combined.svg"
				if *check {
					existing, err := os.ReadFile(svgPath)
					if err != nil || string(existing) != svg {
						fatal(fmt.Errorf("%s: generated chart missing or out of date", svgPath))
					}
				} else {
					if err := os.MkdirAll(filepath.Dir(svgPath), 0o755); err != nil {
						fatal(err)
					}
					if err := os.WriteFile(svgPath, []byte(svg), 0o644); err != nil {
						fatal(err)
					}
				}
			}
		}
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
		{
			charts := "![" + spec.title + " 各库总耗时及输出总大小](charts/" + report.ID + "-combined.svg)\n\n每项仅在全部文件有有效结果时显示；缺失项标为“无完整结果”。大小按表格中的取整数值累加，MB 为 100 万字节。"
			if report.ID == "ofd-to-png" {
				charts += "PNG 大小包含所有页面。"
			}
			charts += "文件较小不一定表示质量更好，请结合还原度等结果查看。"
			for _, metric := range []struct{ id, label string }{{"speed", "耗时"}, {"size", "大小"}} {
				_, missing, err := renderChartSVG(report.Tables, spec, metric.id)
				if err != nil {
					return "", err
				}
				if len(missing) > 0 {
					charts += "\n\n" + metric.label + "无完整结果：" + strings.Join(missing, "；") + "。"
				}
			}
			doc = doc[:start] + "\n\n" + charts + "\n" + doc[end:]
		}
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

func renderCombinedSVG(tables []Table, spec chartSpec) (string, error) {
	type result struct {
		name, color string
		values      [2]float64
		complete    [2]bool
	}
	var sources [2]*Table
	for i := range tables {
		if tables[i].ID == spec.table {
			sources[0] = &tables[i]
		}
		if tables[i].ID == spec.sizeTable {
			sources[1] = &tables[i]
		}
	}
	if sources[0] == nil || sources[1] == nil {
		return "", fmt.Errorf("missing combined chart source")
	}
	colors := []string{"#2563eb", "#0f766e", "#c2410c", "#7e22ce", "#be185d", "#4d7c0f", "#0369a1", "#b45309"}
	var rows []result
	index := map[string]int{}
	var maxima [2]float64
	for metric, table := range sources {
		for _, row := range table.Rows {
			name := row.Label
			if match := markdownLink.FindStringSubmatch(name); match != nil {
				name = match[1]
			}
			i, ok := index[name]
			if !ok {
				i = len(rows)
				index[name] = i
				rows = append(rows, result{name: name, color: colors[i%len(colors)]})
			}
			total := 0.0
			complete := true
			for _, value := range row.Values {
				if value == "FAILED" {
					complete = false
					continue
				}
				n, err := numericValue(value)
				if err != nil {
					return "", fmt.Errorf("%s %s: %w", table.ID, name, err)
				}
				total += n
			}
			if complete {
				divisor := 1000.0
				if metric == 1 {
					divisor = 1000000
				}
				rows[i].values[metric] = total / divisor
				rows[i].complete[metric] = true
				maxima[metric] = math.Max(maxima[metric], total/divisor)
			}
		}
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].complete[0] != rows[j].complete[0] {
			return rows[i].complete[0]
		}
		return rows[i].complete[0] && rows[i].values[0] < rows[j].values[0]
	})
	height := 138 + max(1, len(rows))*54
	var b strings.Builder
	fmt.Fprintf(&b, "<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"1100\" height=\"%d\" viewBox=\"0 0 1100 %d\" role=\"img\" aria-labelledby=\"title desc\">\n", height, height)
	fmt.Fprintf(&b, "<title id=\"title\">%s 总耗时及输出总大小</title>\n<desc id=\"desc\">按总耗时排序，两项各使用独立线性刻度。</desc>\n", html.EscapeString(spec.title))
	fmt.Fprintf(&b, "<rect width=\"1100\" height=\"%d\" rx=\"12\" fill=\"#ffffff\"/>\n", height)
	fmt.Fprintf(&b, "<g font-family=\"Arial, 'Noto Sans CJK SC', sans-serif\" fill=\"#172033\">\n<text x=\"28\" y=\"36\" font-size=\"21\" font-weight=\"700\">%s 总耗时及输出总大小</text>\n<text x=\"28\" y=\"62\" font-size=\"14\" fill=\"#5c6a7d\">%d 个文件合计 · 各栏使用独立刻度</text>\n<text x=\"240\" y=\"95\" font-size=\"15\" font-weight=\"700\">总耗时（秒）</text>\n<text x=\"655\" y=\"95\" font-size=\"15\" font-weight=\"700\">输出总大小（MB）</text>\n", html.EscapeString(spec.title), len(sources[0].Headers)-1)
	for i, row := range rows {
		y := 115 + i*54
		fmt.Fprintf(&b, "<text x=\"28\" y=\"%d\" font-size=\"15\">%s</text>\n", y+18, html.EscapeString(row.name))
		for metric, x := range []int{240, 655} {
			if !row.complete[metric] {
				fmt.Fprintf(&b, "<text x=\"%d\" y=\"%d\" font-size=\"14\" fill=\"#6b7280\">无完整结果</text>\n", x, y+18)
				continue
			}
			width := 0.0
			if maxima[metric] > 0 {
				width = row.values[metric] / maxima[metric] * 280
			}
			fmt.Fprintf(&b, "<rect x=\"%d\" y=\"%d\" width=\"%.2f\" height=\"24\" rx=\"4\" fill=\"%s\"/>\n", x, y, width, row.color)
			fmt.Fprintf(&b, "<text x=\"%.2f\" y=\"%d\" font-size=\"14\" font-weight=\"700\">%.3f</text>\n", float64(x)+width+8, y+18, row.values[metric])
		}
	}
	b.WriteString("</g>\n</svg>\n")
	return b.String(), nil
}

func renderChartSVG(tables []Table, spec chartSpec, metric string) (string, []string, error) {
	tableID, title, divisor, caption, unit, guidance := spec.table, spec.title, 1000.0, "全部文件总耗时", "s", "秒 · 越低越快"
	if metric == "size" {
		tableID, divisor, caption, unit, guidance = spec.sizeTable, 1000000, "输出总大小", "MB", "MB · 越小不一定质量越好"
	}
	type entry struct {
		name  string
		total float64
		color string
	}
	var table *Table
	for i := range tables {
		if tables[i].ID == tableID {
			table = &tables[i]
			break
		}
	}
	if table == nil {
		return "", nil, fmt.Errorf("missing chart source %s", tableID)
	}
	var entries []entry
	var missing []string
	colors := []string{"#2563eb", "#0f766e", "#c2410c", "#7e22ce", "#be185d", "#4d7c0f", "#0369a1", "#b45309"}
	for rowIndex, row := range table.Rows {
		name := row.Label
		if match := markdownLink.FindStringSubmatch(name); match != nil {
			name = match[1]
		}
		var total float64
		var failed []string
		for i, value := range row.Values {
			if value == "FAILED" {
				failed = append(failed, table.Headers[i+1])
				continue
			}
			n, err := numericValue(value)
			if err != nil {
				return "", nil, fmt.Errorf("%s %s %s: %w", tableID, table.Headers[i+1], name, err)
			}
			total += n
		}
		if len(failed) != 0 {
			missing = append(missing, "`"+name+"`（"+strings.Join(failed, "、")+"）")
		} else {
			entries = append(entries, entry{name, total / divisor, colors[rowIndex%len(colors)]})
		}
	}
	sort.SliceStable(entries, func(i, j int) bool { return entries[i].total < entries[j].total })
	height := 120 + max(1, len(entries))*54
	var b strings.Builder
	fmt.Fprintf(&b, "<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"960\" height=\"%d\" viewBox=\"0 0 960 %d\" role=\"img\" aria-labelledby=\"title desc\">\n", height, height)
	fmt.Fprintf(&b, "<title id=\"title\">%s %s</title>\n<desc id=\"desc\">各转换器%s，单位为%s。</desc>\n", html.EscapeString(title), caption, caption, unit)
	fmt.Fprintf(&b, "<rect width=\"960\" height=\"%d\" rx=\"12\" fill=\"#ffffff\"/>\n", height)
	b.WriteString("<g font-family=\"Arial, 'Noto Sans CJK SC', sans-serif\" fill=\"#172033\">\n")
	fmt.Fprintf(&b, "<text x=\"28\" y=\"36\" font-size=\"21\" font-weight=\"700\">%s %s</text>\n<text x=\"28\" y=\"62\" font-size=\"14\" fill=\"#5c6a7d\">%s · %d 个文件</text>\n", html.EscapeString(title), caption, guidance, len(table.Headers)-1)
	maximum := 1.0
	if len(entries) > 0 && entries[len(entries)-1].total > 0 {
		maximum = entries[len(entries)-1].total
	}
	for i, item := range entries {
		y := 96 + i*54
		fmt.Fprintf(&b, "<text x=\"28\" y=\"%d\" font-size=\"15\">%s</text>\n", y+18, html.EscapeString(item.name))
		fmt.Fprintf(&b, "<rect x=\"230\" y=\"%d\" width=\"%.2f\" height=\"24\" rx=\"4\" fill=\"%s\"/>\n", y, item.total/maximum*590, item.color)
		fmt.Fprintf(&b, "<text x=\"%.2f\" y=\"%d\" font-size=\"15\" font-weight=\"700\">%s %s</text>\n", 240+item.total/maximum*590, y+18, strconv.FormatFloat(item.total, 'f', 3, 64), unit)
	}
	if len(entries) == 0 {
		b.WriteString("<text x=\"28\" y=\"114\" font-size=\"15\">暂无完整结果</text>\n")
	}
	b.WriteString("</g>\n</svg>\n")
	return b.String(), missing, nil
}

func updateLibraryVersions(doc *string, reportID string, versions map[string]string) error {
	zc310 := versions["github.com/zc310/ofd"]
	ofdgo := versions["github.com/xiaoqidun/ofdgo"]
	ofdrwRS := versions["github.com/lihongjie0209/ofdrw-rs/bindings/go"]
	ofdrw := versions["org.ofdrw:ofdrw-converter"]
	if reportID == "ofd-to-png" {
		for _, update := range []struct{ prefix, version string }{
			{"| [go-zc310-ofd]", zc310},
			{"| [go-xiaoqidun-ofdgo]", ofdgo},
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
			{"| [go-zc310-ofd]", "✅", 2, zc310},
			{"| [go-xiaoqidun-ofdgo]", "✅", 2, ofdgo},
			{"| [go-ofdrw-rs]", "✅", 2, ofdrwRS},
			{"| [java-ofdrw]", "✅", 2, ofdrw},
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
			{"| [go-zc310-ofd]", "zc310/ofd", 2, zc310},
			{"| [go-xiaoqidun-ofdgo]", "ConvertPDF", 2, ofdgo},
			{"| [java-ofdrw]", "PDFConverter", 2, ofdrw},
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
