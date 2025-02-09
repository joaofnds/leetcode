package main

func isIsomorphic(s string, t string) bool {
	sMap := [256]byte{}
	tMap := [256]byte{}

	for i := range s {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}
		sMap[s[i]] = byte(i + 1)
		tMap[t[i]] = byte(i + 1)
	}

	return true
}
