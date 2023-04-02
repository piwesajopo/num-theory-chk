package main

func fibonacci(n uint) bool {
	nsq5 := 5*n*n
	return perfectSquare(nsq5+4) || perfectSquare(nsq5-4)
}