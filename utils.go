package main

func sumOfDigits(n int) int {
	var s int = 0
	var m, r int
	for r = n; r > 0; r = r / 10 {
		m = r % 10
		s = s + m
	}
	return s
}

func sumOfDigitSquares(n int) int {
	var s int = 0
	var d, r int
	for r = n; r > 0; r = r / 10 {
		d = r % 10
		s = s + d*d
	}
	return s
}
