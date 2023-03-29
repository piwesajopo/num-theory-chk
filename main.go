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

func handleArgs() (int, error) {
	args := os.Args

	debugPtr := flag.Bool("debug", false, "Show Debug Info.")
	flag.Usage = usage
	flag.Parse()

	// Check that at least one parameter has been provided.
	if len(args) < 2 {
		fmt.Println("Sintax: ", args[0], "[-debug] numberToCheck")
		return -1, errors.New("missing parameter")
	}

	params := flag.Args()
	if len(params) == 0 {
		fmt.Println("Sintax: ", args[0], "[-debug] numberToCheck")
		return -1, errors.New("missing parameter")
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
		return -1, errors.New("missing parameter")
	}
	return n, nil
}

func main() {
	number, err := handleArgs()
	if err != nil {
		return
	}

	kind := "happy"
	if debug {
		fmt.Printf("Finding if %d is a %s number.\n", number, kind)
	}

	if happy(number) {
		fmt.Printf("%d IS a %s number.\n", number, kind)
	} else {
		fmt.Printf("%d IS NOT a %s number.\n", number, kind)
	}

	kind = "self"
	if debug {
		fmt.Printf("Finding if %d is a %s number.\n", number, kind)
	}

	if self(number) {
		fmt.Printf("%d IS a %s number.\n", number, kind)
	} else {
		fmt.Printf("%d IS NOT a %s number.\n", number, kind)
	}

}
