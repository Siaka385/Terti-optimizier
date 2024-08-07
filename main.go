package main

import (
	"fmt"
	"os"
	"time"

	"TETRIS-OPTIMIZER/asfunc"
)

func main() {
	start := time.Now()

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

	stop := time.Since(start)

	fmt.Println("Result:", start)
	fmt.Println("Execution time:", stop)
}
