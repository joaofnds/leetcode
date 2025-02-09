package main

func canConstruct(ransomNote string, magazine string) bool {
	counts := [26]int{}
	for _, c := range magazine {
		counts[c-'a']++
	}

	for _, c := range ransomNote {
		counts[c-'a']--
		if counts[c-'a'] < 0 {
			return false
		}
	}

	return true
}
