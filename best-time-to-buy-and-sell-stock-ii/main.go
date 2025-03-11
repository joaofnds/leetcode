package main

func maxProfit(prices []int) int {
	var profit int

	for i := 1; i < len(prices); i++ {
		profit += max(prices[i]-prices[i-1], 0)
	}

	return profit
}
