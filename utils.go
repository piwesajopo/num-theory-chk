package main

import "math"

func sumOfDigits(n uint) uint {
	var s uint = 0
	var m, r uint
	for r = n; r > 0; r = r / 10 {
		m = r % 10
		s = s + m
	}
	return s
}

func sumOfDigitSquares(n uint) uint {
	var s uint = 0
	var d, r uint
	for r = n; r > 0; r = r / 10 {
		d = r % 10
		s = s + d*d
	}
	return s
}

// I was going to make a simple function like others in this file, but couldn't resist to use this
// It is not only faster but provides insight about prime factorization and how it's used to find other divisors
// Credits: github.com/siongui (https://siongui.github.io/2017/06/15/go-calculate-number-of-divisors/)
// Prime factorization of a given number
//
// A map is returned.
//
//	key of map: prime
//	value of map: prime exponents
func PrimeFactorization(n uint) (pfs map[uint]uint) {
	pfs = make(map[uint]uint)

	// Get the number of 2s that divide n
	for n%2 == 0 {
		_, ok := pfs[2]
		if ok {
			pfs[2] += 1
		} else {
			pfs[2] = 1
		}
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := uint(3); i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			_, ok := pfs[i]
			if ok {
				pfs[i] += 1
			} else {
				pfs[i] = 1
			}
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs[n] = 1
	}

	return
}

// Calculate number of divisors of a given number
func NumberOfDivisors(n uint) uint {
	pfs := PrimeFactorization(n)

	var num uint = 1
	for _, exponents := range pfs {
		num *= (exponents + 1)
	}

	return num
}

/*
TODO: Check if this is really a good optimization
      (probably not as good as when used on C)
func IntPow(base, exp int) int {
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}

	return result
}
*/

func IntPow(base, exp uint) uint {
	var result uint = 1
	for {
		if exp%2 == 1 {
			result *= base
		}
		exp /= 2
		if exp == 0 {
			break
		}
		base *= base
	}

	return result
}

func intSquare(n uint) uint {
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

	return result
}

// Adapted from https://siongui.github.io/2017/05/19/go-sum-of-proper-factors/
// formula comes from https://math.stackexchange.com/a/22723
func SumOfProperDivisors(n uint) uint {
	pfs := PrimeFactorization(n)

	var sumOfAllFactors uint = 1
	for prime, exponents := range pfs {
		sumOfAllFactors *= (IntPow(prime, exponents+1) - 1) / (prime - 1)
	}
	return sumOfAllFactors - n
}
