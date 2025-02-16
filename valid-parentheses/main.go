package main

func isValid(s string) bool {
	stack := make([]rune, 0, len(s)/2)

	for _, c := range s {
		closingChar, isOpening := matching(c)
		if isOpening {
			stack = append(stack, closingChar)
		} else if len(stack) > 0 && c == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func matching(r rune) (rune, bool) {
	switch r {
	case '[':
		return ']', true
	case '(':
		return ')', true
	case '{':
		return '}', true
	default:
		return r, false
	}
}
