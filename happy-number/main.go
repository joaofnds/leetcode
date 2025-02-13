package main

func isHappy(n int) bool {
	for n != 1 && n != 4 {
		n = sumOfSquares(n)
	}
	return n == 1
}

func sumOfSquares(n int) int {
	var sum int
	for ; n > 0; n /= 10 {
		sum += (n % 10) * (n % 10)
	}
	return sum
}
