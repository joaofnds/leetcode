package main

func lengthOfLongestSubstring(s string) int {
	var start, longest int
	m := make(map[rune]int, len(s))

	for i, c := range s {
		prevIndex, ok := m[c]
		if ok && prevIndex >= start {
			start = prevIndex + 1
		}
		m[c] = i
		longest = max(longest, i-start+1)
	}

	return longest
}
