package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	a := m - 1
	b := n - 1
	p := m + n - 1

	for a >= 0 && b >= 0 {
		if nums1[a] > nums2[b] {
			nums1[p] = nums1[a]
			a--
		} else {
			nums1[p] = nums2[b]
			b--
		}
		p--
	}

	for b >= 0 {
		nums1[p] = nums2[b]
		b--
		p--
	}
}
