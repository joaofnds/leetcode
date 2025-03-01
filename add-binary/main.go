package main

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	result := make([]byte, max(len(a), len(b))+1)
	r := len(result) - 1
	carry := byte('0')

	for i >= 0 || j >= 0 {
		x := byte('0')
		if i >= 0 {
			x = a[i]
		}

		y := byte('0')
		if j >= 0 {
			y = b[j]
		}

		result[r], carry = fullAdder(x, y, carry)
		i--
		j--
		r--
	}

	if carry == '1' {
		result[r] = carry
	} else {
		result = result[1:]
	}

	return string(result)
}

func fullAdder(a, b, carryIn byte) (byte, byte) {
	sum := a ^ b ^ carryIn
	carryOut := (a & b) | (b & carryIn) | (a & carryIn)
	return sum, carryOut
}
