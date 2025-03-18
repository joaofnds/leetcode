package main

import "strings"

func intToRoman(num int) string {
	var out strings.Builder

	for num > 0 {
		val, symbol := maxRoman(num)
		num -= val
		out.WriteString(symbol)
	}

	return out.String()
}

func maxRoman(num int) (int, string) {
	switch {
	case num >= 1000:
		return 1000, "M"
	case num >= 900:
		return 900, "CM"
	case num >= 500:
		return 500, "D"
	case num >= 400:
		return 400, "CD"
	case num >= 100:
		return 100, "C"
	case num >= 90:
		return 90, "XC"
	case num >= 50:
		return 50, "L"
	case num >= 40:
		return 40, "XL"
	case num >= 10:
		return 10, "X"
	case num >= 9:
		return 9, "IX"
	case num >= 5:
		return 5, "V"
	case num >= 4:
		return 4, "IV"
	default:
		return 1, "I"
	}
}
