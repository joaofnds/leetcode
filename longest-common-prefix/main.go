package main

func longestCommonPrefix(strs []string) string {
	for i := range strs[0] {
		for _, str := range strs[1:] {
			if i >= len(str) || str[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
