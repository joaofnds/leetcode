package main

func reverseBits(num uint32) uint32 {
	var res uint32
	for range 32 {
		res <<= 1
		res |= num & 1
		num >>= 1
	}
	return res
}
