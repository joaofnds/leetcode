package main

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens)/2)

	for _, token := range tokens {
		n := len(stack)
		switch token {
		case "+":
			stack = append(stack[:n-2], stack[n-2]+stack[n-1])
		case "-":
			stack = append(stack[:n-2], stack[n-2]-stack[n-1])
		case "*":
			stack = append(stack[:n-2], stack[n-2]*stack[n-1])
		case "/":
			stack = append(stack[:n-2], stack[n-2]/stack[n-1])
		default:
			stack = append(stack, atoi(token))
		}
	}

	return stack[0]
}

func atoi(s string) int {
	sign, num := 1, 0

	if s[0] == '-' {
		sign = -1
	}

	for _, c := range s {
		if '0' <= c && c <= '9' {
			num = num*10 + int(c-'0')
		}
	}

	return sign * num
}
