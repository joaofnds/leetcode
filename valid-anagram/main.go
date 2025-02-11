package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := [26]int{}

	for i := range s {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
