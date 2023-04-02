package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var debug bool

func usage() {
	/*
			Original Message:
				flag provided but not defined: -t
				Usage of ./self-number:
		  		  -debug
		    			Show Debug Info.
	*/
	fmt.Fprintf(os.Stderr, "Sintax: %s [-debug] numberToCheck\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func handleArgs() (uint, error) {
	args := os.Args

	debugPtr := flag.Bool("debug", false, "Show Debug Info.")
	flag.Usage = usage
	flag.Parse()

	// Check that at least one parameter has been provided.
	if len(args) < 2 {
		fmt.Println("Sintax: ", args[0], "[-debug] numberToCheck")
		return 0, errors.New("missing parameter")
	}

	params := flag.Args()
	if len(params) == 0 {
		fmt.Println("Sintax: ", args[0], "[-debug] numberToCheck")
		return 0, errors.New("missing parameter")
	}

	// Use flag.CommandLine.Parse to parse the remaining arguments.
	// This will cover the case when the -debug flag was specified after the number
	flag.CommandLine.Parse(flag.Args()[1:])

	// Obtain value of -debug parameter (true if parameter provided)
	debug = *debugPtr

	n, err := strconv.Atoi(params[0])
	if err != nil {
		fmt.Println("Please specify a number.")
		fmt.Println("Sintax: ", args[0], "[-debug] numberToCheck")
		return 0, errors.New("missing parameter")
	}
	return uint(n), nil
}

type evalKind func(n uint) bool

func testKind(kind string, number uint, fn evalKind) {
	var article string
	switch kind[0] {
	case
		'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		article = "an"
	default:
		article = "a"
	}

	if debug {
		fmt.Printf("Finding if %d is %s %s number.\n", number, article, kind)
	}

	if fn(number) {
		fmt.Printf("%d IS %s %s number.\n", number, article, kind)
	} else {
		if debug {
			fmt.Printf("%d IS NOT %s %s number.\n", number, article, kind)
		}
	}
}

func main() {
	number, err := handleArgs()
	if err != nil {
		return
	}

	sumOfProperDiv := SumOfProperDivisors(number)
	fmt.Println("Number", number, "has", NumberOfDivisors(number), "divisors.")
	fmt.Println("Sum of proper divisors of", number, "is", sumOfProperDiv)

	if number == sumOfProperDiv {
		fmt.Println("Number", number, "is a perfect number")
	} else if number < sumOfProperDiv {
		fmt.Println("Number", number,
			"is an abundant number with an abundance of", sumOfProperDiv-number)
	} else {
		fmt.Println("Number", number,
			"is an deficient number with an deficiency of", number-sumOfProperDiv)
	}

	testKind("perfect square", number, perfectSquare)
	testKind("happy", number, happy)
	testKind("self", number, self)
	testKind("odious", number, odious)
}
