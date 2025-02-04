package main

var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue := romanMap[s[i]]

		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}

		prevValue = currentValue
	}

	return result
}
