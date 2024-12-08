package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	levels := [][]int{}
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text() // Read the current line as a string

		parts := strings.Split(line, " ")

		level := []int{}
		for _, valStr := range parts {
			val, _ := strconv.Atoi(valStr)
			level = append(level, val)
		}
		levels = append(levels, level)
	}

	safeReportCount := countSafeReports(levels)

	fmt.Printf("Result: %d\n", safeReportCount)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func countSafeReports(levels [][]int) int {
	var count = 0

	for _, level := range levels {
		isSortedFunny := true
		isAsc := false
		isDesc := false
		for i := len(level) - 1; i > 0; i-- {
			if level[i] < level[i-1] && level[i]+3 >= level[i-1] {
				isAsc = true
			} else if level[i] > level[i-1] && level[i] <= level[i-1]+3 {
				isDesc = true
			} else {
				isSortedFunny = false
				break
			}

			if isAsc && isDesc {
				isSortedFunny = false
				break
			}
		}

		if isSortedFunny {
			count += 1
		}
	}

	return count
}
