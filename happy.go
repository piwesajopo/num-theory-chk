package main

import "fmt"

func happy(n uint) bool {
	for i := 0; ; i++ {
		n = sumOfDigitSquares(n)
		if debug {
			fmt.Println("Sum of Digits Squares:", n)
		}
		if n == 1 {
			return true
		}
		if n == 4 {
			return false
		}
	}
}
