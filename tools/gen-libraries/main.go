package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"

	"go.yaml.in/yaml/v3"
)

type Column struct {
	Key        string `yaml:"key"`
	Header     string `yaml:"header"`
	Legend     string `yaml:"legend,omitempty"`
	Structural bool   `yaml:"structural,omitempty"`
}

type Section struct {
	Lang  string `yaml:"lang"`
	Title string `yaml:"title"`
}

type Library struct {
	Name    string            `yaml:"name"`
	URL     string            `yaml:"url"`
	Lang    string            `yaml:"lang"`
	License string            `yaml:"license"`
	Caps    map[string]string `yaml:"caps"`
}

type Document struct {
	Title     string    `yaml:"title"`
	Sections  []Section `yaml:"sections"`
	Columns   []Column  `yaml:"columns"`
	Libraries []Library `yaml:"libraries"`
}

var errCheck = errors.New("output out of date")

func main() {
	src := flag.String("src", "docs/libraries.yaml", "数据源 YAML")
	out := flag.String("out", "docs/ofd-libraries.md", "生成目标 Markdown")
	check := flag.Bool("check", false, "仅校验生成结果是否与磁盘一致")
	flag.Parse()

	raw, err := os.ReadFile(*src)
	if err != nil {
		fatal(err)
	}
	var doc Document
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		fatal(fmt.Errorf("%s: %w", *src, err))
	}
	if err := validate(doc); err != nil {
		fatal(fmt.Errorf("%s: %w", *src, err))
	}
	rendered := render(doc)

	if *check {
		cur, err := os.ReadFile(*out)
		if err != nil {
			fatal(err)
		}
		if string(cur) != rendered {
			fatal(fmt.Errorf("%s: %w（运行 go run ./tools/gen-libraries 重新生成）", *out, errCheck))
		}
		return
	}
	if err := os.WriteFile(*out, []byte(rendered), 0o644); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "gen-libraries:", err)
	os.Exit(1)
}

func validate(doc Document) error {
	if len(doc.Title) == 0 {
		return errors.New("title 为空")
	}
	seenSection := map[string]bool{}
	for _, s := range doc.Sections {
		if s.Lang == "" || s.Title == "" {
			return errors.New("section 缺少 lang/title")
		}
		if seenSection[s.Lang] {
			return fmt.Errorf("section 重复: %s", s.Lang)
		}
		seenSection[s.Lang] = true
	}
	seenCol := map[string]bool{}
	var capKeys []string
	for _, c := range doc.Columns {
		if c.Key == "" || c.Header == "" {
			return errors.New("column 缺少 key/header")
		}
		if seenCol[c.Key] {
			return fmt.Errorf("column 重复: %s", c.Key)
		}
		seenCol[c.Key] = true
		if !c.Structural {
			capKeys = append(capKeys, c.Key)
		}
	}
	seenURL := map[string]bool{}
	for i, lib := range doc.Libraries {
		if lib.Name == "" || lib.URL == "" || lib.License == "" {
			return fmt.Errorf("第 %d 条库缺少 name/url/license", i+1)
		}
		if seenURL[lib.URL] {
			return fmt.Errorf("url 重复: %s", lib.URL)
		}
		seenURL[lib.URL] = true
		if !seenSection[lib.Lang] {
			return fmt.Errorf("%s: 未知 section %q", lib.Name, lib.Lang)
		}
		for _, k := range capKeys {
			v, ok := lib.Caps[k]
			if !ok {
				return fmt.Errorf("%s: 缺少能力 %s", lib.Name, k)
			}
			if v != "✅" && v != "⚠️" && v != "❌" {
				return fmt.Errorf("%s: %s 取值非法 %q", lib.Name, k, v)
			}
		}
		for k := range lib.Caps {
			if !seenCol[k] {
				return fmt.Errorf("%s: 未定义列 %s", lib.Name, k)
			}
		}
	}
	for _, s := range doc.Sections {
		found := false
		for _, lib := range doc.Libraries {
			if lib.Lang == s.Lang {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("section %s 无库条目", s.Title)
		}
	}
	return nil
}

func render(doc Document) string {
	var b strings.Builder
	b.WriteString("<!-- 本文件由 tools/gen-libraries 生成，勿直接编辑；数据源 docs/libraries.yaml -->\n")
	b.WriteString("# " + doc.Title + "\n")

	var legendLines []string
	for _, c := range doc.Columns {
		if c.Legend != "" {
			legendLines = append(legendLines, "> **"+c.Header+"**："+c.Legend)
		}
	}
	for i, l := range legendLines {
		if i > 0 {
			b.WriteString("> \n")
		}
		b.WriteString(l + "\n")
	}
	b.WriteString("\n")

	for _, sec := range doc.Sections {
		b.WriteString("## " + sec.Title + "\n\n")
		b.WriteString(renderTable(doc, sec.Lang))
		b.WriteString("\n")
	}
	return b.String()
}

func renderTable(doc Document, lang string) string {
	var libs []Library
	for _, lib := range doc.Libraries {
		if lib.Lang == lang {
			libs = append(libs, lib)
		}
	}
	n := len(doc.Columns)
	widths := make([]int, n)
	for i, c := range doc.Columns {
		widths[i] = dispWidth(c.Header) + 2
	}
	for _, lib := range libs {
		for i, c := range doc.Columns {
			w := dispWidth(cellContent(c, lib)) + 2
			if w > widths[i] {
				widths[i] = w
			}
		}
	}
	for i := range widths {
		if widths[i] < 5 {
			widths[i] = 5
		}
	}

	var b strings.Builder
	writeRow(&b, widths, func(i int) string { return doc.Columns[i].Header })
	writeSep(&b, widths)
	for _, lib := range libs {
		lib := lib
		writeRow(&b, widths, func(i int) string { return cellContent(doc.Columns[i], lib) })
	}
	return b.String()
}

func cellContent(c Column, lib Library) string {
	switch c.Key {
	case "name":
		return "[" + lib.Name + "](" + lib.URL + ")"
	case "license":
		return lib.License
	default:
		return lib.Caps[c.Key]
	}
}

func writeRow(b *strings.Builder, widths []int, content func(int) string) {
	b.WriteString("|")
	for i, w := range widths {
		b.WriteString(center(content(i), w))
		b.WriteString("|")
	}
	b.WriteString("\n")
}

func writeSep(b *strings.Builder, widths []int) {
	b.WriteString("|")
	for _, w := range widths {
		b.WriteString(":" + strings.Repeat("-", w-2) + ":")
		b.WriteString("|")
	}
	b.WriteString("\n")
}

func center(s string, w int) string {
	pad := w - dispWidth(s)
	if pad < 0 {
		pad = 0
	}
	left := pad / 2
	right := pad - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}

func dispWidth(s string) int {
	total := 0
	for _, r := range s {
		total += runeWidth(r)
	}
	return total
}

func runeWidth(r rune) int {
	switch {
	case r == '\u200d' || r == '\ufe0f':
		return 0
	case r == '\u2705' || r == '\u274c' || r == '\u26a0':
		return 2
	case isWide(r):
		return 2
	default:
		return 1
	}
}

func isWide(r rune) bool {
	switch {
	case r >= 0x1100 && r <= 0x115f,
		r >= 0x2e80 && r <= 0x303e,
		r >= 0x3041 && r <= 0x33ff,
		r >= 0x3400 && r <= 0x4dbf,
		r >= 0x4e00 && r <= 0x9fff,
		r >= 0xa000 && r <= 0xa4cf,
		r >= 0xac00 && r <= 0xd7a3,
		r >= 0xf900 && r <= 0xfaff,
		r >= 0xfe30 && r <= 0xfe4f,
		r >= 0xff00 && r <= 0xff60,
		r >= 0xffe0 && r <= 0xffe6,
		r >= 0x1f300 && r <= 0x1faff,
		r >= 0x20000 && r <= 0x3fffd:
		return true
	case unicode.Is(unicode.Mn, r), unicode.Is(unicode.Me, r):
		return false
	default:
		return false
	}
}
