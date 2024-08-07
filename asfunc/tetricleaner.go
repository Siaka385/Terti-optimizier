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
	var validLines [][]string
	columnHasHash := make(map[int]bool)

	// Split each line into characters and collect valid lines and columns to remove
	for _, line := range lines {
		chars := strings.Split(line, "")
		hasHash := false
		for j, char := range chars {
			if char == "#" {
				hasHash = true
				columnHasHash[j] = true
			}
		}
		if hasHash {
			validLines = append(validLines, chars)
		}
	}

	// Remove columns without any hash
	var cleanedLines [][]string
	for _, line := range validLines {
		var cleanedLine []string
		for j, char := range line {
			if columnHasHash[j] {
				cleanedLine = append(cleanedLine, char)
			}
		}
		cleanedLines = append(cleanedLines, cleanedLine)
	}

	// Convert cleaned 2D slice back to string
	var result []string
	for _, line := range cleanedLines {
		result = append(result, strings.Join(line, ""))
	}

	return strings.Join(result, "\n")
}
