package main

func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}

	buckets := make([]int, n+1)

	for _, c := range citations {
		buckets[min(c, n)]++
	}

	count := 0
	for i := n; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}

	return 0
}
