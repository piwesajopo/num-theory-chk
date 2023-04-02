package main

import (
	"math"
)

// Based on https://www.geeksforgeeks.org/find-square-root-of-a-number-using-bit-manipulation/
func perfectSquare(n uint) bool {
	// Find MSB (Most significant Bit) of N
	msb := uint(math.Log2(float64(n)))

	// (a = 2^msb)
	var a uint = 1 << msb
	var result uint = 0
	for a != 0 {
		// Check whether the current value
		// of 'a' can be added or not
		if (result+a)*(result+a) <= n {
			result += a
		}

		// (a = a/2)
		a >>= 1
	}

	return result*result == n
}
