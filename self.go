package main

import (
	"fmt"
	"math"
)

func self(n uint) bool {

	// Actually the number of digits - 1 (but that's what we want)
	digits := uint(math.Floor(math.Log10(float64(n))))

	var maxD uint = digits * 9

	lowerBoundary := n - maxD

	var checkX uint
	for number := uint(lowerBoundary); number < n; number++ {
		checkX = number + sumOfDigits(number)
		if debug {
			fmt.Printf("Checking for %d. Generated: %d\n", number, checkX)
		}
		if checkX >= n {
			break
		}
	}
	return n != checkX
}
