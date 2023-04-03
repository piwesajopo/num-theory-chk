package main

import "math"

func armstrong(n uint) (isNarcissistic bool, power uint) {
	// Digit count of a base10 number = floor(log10(number)+1)
	var c uint = uint(math.Floor(math.Log10(float64(n)) + 1)) // Floor(log10(n) + 1)
	var s uint = 0

	var d, r uint
	for r = n; r > 0; r = r / 10 {
		d = r % 10
		s = s + IntPow(d, c)
	}

	isNarcissistic = s == n
	power = c 
	return
}
