package main

func lengthOfLastWord(s string) int {
	l := 0

	for i := len(s) - 1; i >= 0; i-- {
		if 'A' <= s[i] && s[i] <= 'z' {
			l++
		} else if l > 0 {
			return l
		}
	}

	return l
}
