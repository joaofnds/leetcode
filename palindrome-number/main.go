package main

func isPalindrome(x int) bool {
	var reversed int
	for n := x; n > 0; n /= 10 {
		reversed = reversed*10 + n%10
	}
	return x == reversed
}
