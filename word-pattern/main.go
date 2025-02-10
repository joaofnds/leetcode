package main

func wordPattern(pattern string, s string) bool {
	var p, i, j int
	m := map[string]byte{}
	used := [26]bool{}

	for p < len(pattern) && i < len(s) {
		if j < len(s) && s[j] != ' ' {
			j++
			continue
		}

		if s[i] == ' ' {
			i++
			continue
		}

		char := pattern[p]
		word := s[i:j]
		charIndex := char - 'a'

		existing, ok := m[word]
		if ok && existing != char || !ok && used[charIndex] {
			return false
		}

		m[word] = char
		used[charIndex] = true

		p++
		j++
		i = j
	}

	return p == len(pattern) && j > len(s)
}
