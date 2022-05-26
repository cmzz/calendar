package calendar

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func PrintTable(headers []string, rows [][]string, footers []string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	addHeader(t, headers)
	addRows(t, rows)
	t.AppendSeparator()
	addFooter(t, footers)
	t.Render()
}

func addHeader(t table.Writer, headers []string) {
	if headers == nil {
		return
	}

	r := table.Row{}
	for _, item := range headers {
		r = append(r, item)
	}

	t.AppendHeader(r)
}

func addRows(t table.Writer, rows [][]string) {
	if rows == nil {
		return
	}

	var tb []table.Row
	for _, row := range rows {
		r := table.Row{}
		for _, cell := range row {
			r = append(r, cell)
		}
		tb = append(tb, r)
	}

	t.AppendRows(tb)
}

func addFooter(t table.Writer, footers []string) {
	if footers == nil {
		return
	}
	r := table.Row{}
	for _, item := range footers {
		r = append(r, item)
	}
	t.AppendFooter(r)
}
