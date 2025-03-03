package main

func hammingWeight(n int) int {
	var count int
	for ; n > 0; n >>= 1 {
		count += n & 1
	}
	return count
}
