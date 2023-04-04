package main

// These functions are adapted from: https://www.geeksforgeeks.org/semiperfect-number/

func sliceFactors(number uint) (factors []int) {
	n := int(number)
	sqrt := int(intSquare(number))
	factors = append(factors, 1)

	// note that this loop runs till sqrt(n)
	for i := 2; i <= sqrt; i++ {
		// if the value of i is a factor
		if n%i == 0 {
			factors = append(factors, i)

			// condition to check the
			// divisor is not the number itself
			if n/i != i {
				factors = append(factors, n/i)
			}
		}
	}
	// return the vector
	return
}

func semiperfect(n uint) bool {
	// This returns a slice of all the proper divisors
	// Note: The returned slice is already sorted, which
	//       is a pre-requisite for this code to work
	//factors := slicePrimeFactors(n)
	factors := sliceFactors(n)

	r := len(factors) + 1
	c := int(n) + 1

	// Creates Slice [r][c]bool
	// Note: all values are initialized to false
	subset := make([][]bool, r)
	for i := 0; i < r; i++ {
		subset[i] = make([]bool, c)
	}

	// Let's initializa the first element of each column to true
	for i := 0; i < r; i++ {
		subset[i][0] = true
	}

	// loop to find whether the
	// number is semiperfect
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			// calculation to check if the
			// number can be made by
			// summation of divisors
			if j < factors[i-1] {
				subset[i][j] = subset[i-1][j]
			} else {
				subset[i][j] = subset[i-1][j] || subset[i-1][j-factors[i-1]]
			}
		}
	}

	// If not possible to make the
	// number by any combination of divisors
	return subset[r-1][c-1]
}
