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
		fmt.Printf("Finding if %d is %s %s number\n", number, article, kind)
	}

	if fn(number) {
		fmt.Printf("%d is %s %s number\n", number, article, kind)
	} else {
		if debug {
			fmt.Printf("%d IS NOT %s %s number\n", number, article, kind)
		}
	}
}

func main() {
	number, err := handleArgs()
	if err != nil {
		return
	}

	sumOfProperDiv := sumOfProperDivisors(number)
	numDiv := numberOfDivisors(number)
	if numDiv == 2 {
		fmt.Println(number, "is a prime number")
	} else {
		fmt.Println(number, "has", numDiv, "divisors")
		fmt.Println("Sum of proper divisors of", number, "is", sumOfProperDiv)
	}

	if number == sumOfProperDiv {
		fmt.Println(number, "is a perfect number")
	} else if number < sumOfProperDiv {
		fmt.Println(number,
			"is an abundant number with an abundance of", sumOfProperDiv-number)
	} else {
		fmt.Println(number,
			"is an deficient number with an deficiency of", number-sumOfProperDiv)
	}

	if debug {
		fmt.Printf("Finding if %d is a Narcissistic number\n", number)
	}

	isNarcissistic, n := armstrong(number)
	if isNarcissistic {
		fmt.Printf("%d is a %d-Narcissistic number\n", number, n)
	} else {
		if debug {
			fmt.Printf("%d IS NOT a %d-Narcissistic number\n", number, n)
		}
	}

	testKind("perfect square", number, perfectSquare)
	testKind("happy", number, happy)
	testKind("self", number, self)
	testKind("odious", number, odious)
	testKind("fibonacci", number, fibonacci)
	testKind("münchhausen", number, munchhausen)
	testKind("factorial", number, factorial)

	if debug {
		fmt.Printf("Finding if %d is a Narcissistic number\n", number)
	}
	isSemiperfect, err := semiperfect(number)
	if err != nil {
		fmt.Println(number, err)
		return
	}

	if isSemiperfect {
		fmt.Printf("%d is a semiperfect number\n", number)
	} else {
		if debug {
			fmt.Printf("%d IS NOT a semiperfect number\n", number)
		}
	}
}
