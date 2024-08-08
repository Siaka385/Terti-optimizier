package asfunc

import (
	"strings"
)

// Function to clean each chunk using the janitor function
func Cleaner(tetri []string) []string {
	var cleanedtetri []string
	for _, chunk := range tetri {
		cleanedtetri = append(cleanedtetri, janitor(chunk))
	}
	return cleanedtetri
}

// Function to clean a single chunk

func janitor(chunk string) string {

	lines := strings.Split(chunk, "\n")
	var cleanedLines [][]string
	columnHasHash := make(map[int]bool)

	for _, line := range lines {
		chars := strings.Split(line, "")
		hasHash := false
		var cleanedLine []string

		for j, char := range chars {
			if char == "#" {
				hasHash = true
				columnHasHash[j] = true
			}
			if columnHasHash[j] {
				cleanedLine = append(cleanedLine, char)
			}
		}

		if hasHash {
			cleanedLines = append(cleanedLines, cleanedLine)
		}
	}

	var result []string
	for _, line := range cleanedLines {
		result = append(result, strings.Join(line, ""))
	}

	return strings.Join(result, "\n")
}
