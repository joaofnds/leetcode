package main

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]uint8][]string)

	for _, str := range strs {
		count := [26]uint8{}
		for _, char := range str {
			count[char-'a']++
		}
		anagrams[count] = append(anagrams[count], str)
	}

	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}
