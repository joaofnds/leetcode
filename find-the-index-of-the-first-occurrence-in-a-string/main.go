package main

func strStr(haystack string, needle string) int {
	var h, n int

	for h < len(haystack) {
		if haystack[h] == needle[n] {
			h++
			n++
			if n == len(needle) {
				return h - n
			}
		} else {
			h = h - n + 1
			n = 0
		}
	}

	return -1
}
