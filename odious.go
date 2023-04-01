package main

// Returns true if the n is an odious number
func odious(n uint) bool {
	var cnt uint
	// Notice we use n >> 1 here instead of n / 2
	// It may seem it was done for performance
	// but it's because it makes more sense given
	// the purpose of this function which deals with
	// binary data.
	for cnt = 0; n > 0; n >>= 1 {
		// Here we also use the & operation in n & 1'
		// instead of using n % 2 which also serves better
		// as it self documents the intention of the code.
		lsbit := n & 1 // turns off all bits but the least significant
		cnt += lsbit
	}

	// Notice here we use 'cnt % 2' even when 'cnt & 1 would serve the same
	// purpose (implicit documentation is more important than performance).
	return cnt%2 == 1 // If cnt if odd then number is odious
}
