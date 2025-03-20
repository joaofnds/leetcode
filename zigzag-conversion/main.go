package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	period := 2*numRows - 2

	for i := range s {
		mod := i % period
		rows[min(mod, period-mod)].WriteByte(s[i])
	}

	var out strings.Builder
	out.Grow(len(s))
	for _, row := range rows {
		out.WriteString(row.String())
	}
	return out.String()
}
