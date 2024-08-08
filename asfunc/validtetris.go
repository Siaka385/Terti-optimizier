package asfunc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Collect(filename string) (bool, []string) {
	if !strings.HasSuffix(filename, ".txt") {
		return false, []string{}
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	var mytetris []string

	mytetris = append(mytetris, " \n")
	line := ""
	// Scan the file line by line
	for scanner.Scan() {
		counter++
		// Get the current line
		line += strings.TrimSpace(scanner.Text()) + "\n"

		if counter > 4 {
			if scanner.Text() != "" {
				return false, mytetris
			}
			mytetris = append(mytetris, line[:len(line)-2])
			counter = 0
			line = ""

		}
	}
	if line != "" {
		mytetris = append(mytetris, line[:len(line)-1])
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return true, mytetris
}

func IsconnectedAtJoints(m string) bool {
	mystring := strings.Split(m, "\n")
	if len(mystring) != 4 {
		return false
	}
	for _, line := range mystring {
		if len(line) != 4 {
			return false
		}
	}

	// Convert mystring slice to a 2D slice of strings
	var mymothertetris [][]string
	for _, line := range mystring {
		mymothertetris = append(mymothertetris, strings.Split(line, ""))
	}

	connections := 0
	hashCount := 0

	for i := 0; i < len(mymothertetris); i++ {
		for k := 0; k < len(mymothertetris[i]); k++ {
			if mymothertetris[i][k] == "#" {
				hashCount++
				if i == 0 {
					if k == 0 {
						if mymothertetris[i][k+1] == "#" {
							connections++
						}
						if mymothertetris[i+1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k+1] != "#" && mymothertetris[i+1][k] != "#" {
							return false
						}
					} else if k == 3 {
						if mymothertetris[i][k-1] == "#" {
							connections++
						}
						if mymothertetris[i+1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k-1] != "#" && mymothertetris[i+1][k] != "#" {
							return false
						}
					} else {
						if mymothertetris[i][k-1] == "#" {
							connections++
						}
						if mymothertetris[i+1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k+1] == "#" {
							connections++
						}
						if mymothertetris[i][k-1] != "#" && mymothertetris[i+1][k] != "#" && mymothertetris[i][k+1] != "#" {
							return false
						}
					}
				} else if i == 3 {
					if k == 0 {
						if mymothertetris[i][k+1] == "#" {
							connections++
						}
						if mymothertetris[i-1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k+1] != "#" && mymothertetris[i-1][k] != "#" {
							return false
						}
					} else if k == 3 {
						if mymothertetris[i][k-1] == "#" {
							connections++
						}
						if mymothertetris[i-1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k-1] != "#" && mymothertetris[i-1][k] != "#" {
							return false
						}
					} else {
						if mymothertetris[i][k-1] == "#" {
							connections++
						}
						if mymothertetris[i-1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k+1] == "#" {
							connections++
						}
						if mymothertetris[i][k-1] != "#" && mymothertetris[i-1][k] != "#" && mymothertetris[i][k+1] != "#" {
							return false
						}
					}
				} else {
					if k == 0 {
						if mymothertetris[i][k+1] == "#" {
							connections++
						}
						if mymothertetris[i-1][k] == "#" {
							connections++
						}
						if mymothertetris[i+1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k+1] != "#" && mymothertetris[i-1][k] != "#" && mymothertetris[i+1][k] != "#" {
							return false
						}
					} else if k == 3 {
						if mymothertetris[i][k-1] == "#" {
							connections++
						}
						if mymothertetris[i-1][k] == "#" {
							connections++
						}
						if mymothertetris[i+1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k-1] != "#" && mymothertetris[i-1][k] != "#" && mymothertetris[i+1][k] != "#" {
							return false
						}
					} else {
						if mymothertetris[i][k-1] == "#" {
							connections++
						}
						if mymothertetris[i-1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k+1] == "#" {
							connections++
						}
						if mymothertetris[i+1][k] == "#" {
							connections++
						}
						if mymothertetris[i][k-1] != "#" && mymothertetris[i-1][k] != "#" && mymothertetris[i][k+1] != "#" && mymothertetris[i+1][k] != "#" {
							return false
						}
					}
				}
			}
		}
	}

	if connections < 6 || connections > 8 {
		return false
	}
	if hashCount != 4 {
		return false
	}

	return true
}
