package report

import (
	"bytes"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// MarkdownTable returns a text-formatted table with caption, optionally as valid Markdown,
// so the output can be pasted into the docs.
func MarkdownTable(rows [][]string, cols []string, caption string, valid bool) string {
	// Escape Markdown.
	if valid {
		for i := range rows {
			for j := range rows[i] {
				if strings.ContainsRune(rows[i][j], '|') {
					rows[i][j] = strings.ReplaceAll(rows[i][j], "|", "\\|")
				}
			}
		}
	}

	buf := &bytes.Buffer{}

	// Set Borders.
	borders := tablewriter.Border{
		Left:   true,
		Right:  true,
		Top:    !valid,
		Bottom: !valid,
	}

	// Render.
	table := tablewriter.NewWriter(buf)

	// Set Caption.
	if caption != "" {
		table.SetCaption(true, caption)
	}

	table.SetAutoWrapText(!valid)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(cols)
	table.SetBorders(borders)
	table.SetCenterSeparator("|")
	table.AppendBulk(rows)
	table.Render()

	return buf.String()
}
