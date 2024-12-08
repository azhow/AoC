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
		isSortedFunny, idxNotFunny := checkSortedFunny(level)

		if isSortedFunny {
			count += 1
		} else {

			levelRight := append(level[:idxNotFunny:idxNotFunny], level[idxNotFunny+1:]...)
			isSortedFunnyRight, _ := checkSortedFunny(levelRight)
			levelLeft := append(level[:idxNotFunny-1:idxNotFunny-1], level[idxNotFunny:]...)
			isSortedFunnyLeft, _ := checkSortedFunny(levelLeft)

			if isSortedFunnyLeft {
				count += 1
			} else if isSortedFunnyRight {
				count += 1
			} else {
				fmt.Printf("%#v\n", level)
			}
		}
	}

	return count
}

func checkSortedFunny(level []int) (bool, int) {
	isSortedFunny := true
	isAsc := false
	isDesc := false
	count := len(level) - 1
	for i := 1; i < len(level); i++ {
		if level[i] < level[i-1] && level[i]+3 >= level[i-1] {
			isDesc = true
		} else if level[i] > level[i-1] && level[i] <= level[i-1]+3 {
			isAsc = true
		} else {
			isSortedFunny = false
			count = i
			break
		}

		if isAsc && isDesc {
			isSortedFunny = false
			count = i
			break
		}
	}

	return isSortedFunny, count
}
