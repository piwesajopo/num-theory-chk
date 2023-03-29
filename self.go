package main

import (
	"fmt"
	"math"
)

func self(n int) bool {

	// Actually the number of digits - 1 (but that's what we want)
	digits := int(math.Floor(math.Log10(float64(n))))

	maxD := digits * 9

	lowerBoundary := n - maxD

	var checkX int
	for number := lowerBoundary; number < n; number++ {
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
