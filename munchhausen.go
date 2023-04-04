package main

// This assumes 0^0 = 1, thus only valid # is 3435
func munchhausen(n uint) bool {
	var s uint = 0
	var d, r uint
	for r = n; r > 0; r = r / 10 {
		d = r % 10
		s = s + intPow(d, d)
	}
	return s == n
}
