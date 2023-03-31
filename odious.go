package main

func odious(n int) bool {
	var cnt int
	for cnt = 0; n > 0; n >>= 1 {
		// Least significant bit is 1 when number is not even (i.e. n%2 == 1)
		lsbit := n % 2
		cnt += lsbit
	}

	return cnt%2 == 1 // If cnt if odd then number is odious
}
