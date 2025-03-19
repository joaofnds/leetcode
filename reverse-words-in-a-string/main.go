package main

import "strings"

func reverseWords(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}

		j := i
		for j >= 0 && s[j] != ' ' {
			j--
		}

		b.WriteByte(' ')
		b.WriteString(s[j+1 : i+1])

		i = j
	}

	return b.String()[1:]
}
