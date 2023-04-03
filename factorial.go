package main

func factorial(n uint) bool {

	var mod, mult uint
	for mult = 1; ; mult++ {
		mod = n % mult
		if mod == 0 {
			n /= mult
		} else {
			break
		}
	}

	// if all were divisors number is the result of a factorial
	return n == 1
}