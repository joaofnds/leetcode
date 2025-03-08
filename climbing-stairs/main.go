package main

func climbStairs(n int) int {
	first, second := 1, 1
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}
