package main

import (
	"fmt"
	"os"

	"TETRIS-OPTIMIZER/asfunc"
)

/**
 * main is the entry point of the program.
 * It collects the input arguments, validates them, and performs Tetris-related operations.
 * If any errors occur during the process, it prints "ERROR" and exits with status code 1.
 */

func main() {
	myargs := os.Args[1:]
	if len(myargs) != 1 {
		fmt.Println("ERROR")
		os.Exit(1)
	}
	valid, tetri := asfunc.Collect(myargs[0])

	if valid {
		tetri = tetri[1:]
		for i := 0; i < len(tetri); i++ {
			if !asfunc.IsconnectedAtJoints(tetri[i]) {
				fmt.Println("ERROR")
				os.Exit(1)
			}
		}

		tetri = asfunc.Cleaner(tetri)

		asfunc.Proo(tetri, len(tetri))
	} else {
		fmt.Println("ERROR")
		os.Exit(1)
	}
}
