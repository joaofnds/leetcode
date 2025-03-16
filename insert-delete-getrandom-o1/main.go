package main

import "math/rand"

type RandomizedSet struct {
	nums []int
	pos  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		pos:  map[int]int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.pos[val]; exists {
		return false
	}

	rs.nums = append(rs.nums, val)
	rs.pos[val] = len(rs.nums) - 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	pos, exists := rs.pos[val]
	if !exists {
		return false
	}

	if last := rs.nums[len(rs.nums)-1]; val != last {
		rs.nums[pos], rs.pos[last] = last, pos
	}

	rs.nums = rs.nums[:len(rs.nums)-1]
	delete(rs.pos, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	if len(rs.nums) == 0 {
		return 0
	}

	return rs.nums[rand.Intn(len(rs.nums))]
}
