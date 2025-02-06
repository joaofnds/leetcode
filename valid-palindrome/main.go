package main

import "strings"

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	lower := strings.ToLower(s)

	for i < j {
		if !isAlphanumeric(lower[i]) {
			i++
			continue
		}

		if !isAlphanumeric(lower[j]) {
			j--
			continue
		}

		if lower[i] != lower[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	return ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')
}
