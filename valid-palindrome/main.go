package main

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		left, right := toLower(s[i]), toLower(s[j])

		if !isAlphanumeric(left) {
			i++
			continue
		}

		if !isAlphanumeric(right) {
			j--
			continue
		}

		if left != right {
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

func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + ('a' - 'A')
	}

	return c
}
