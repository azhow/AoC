package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	leftVals := []float64{}
	rightVals := []float64{}

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text() // Read the current line as a string

		parts := strings.Split(line, "   ")
		//fmt.Println(parts)
		left, _ := strconv.ParseFloat(parts[0], 64)
		right, _ := strconv.ParseFloat(parts[1], 64)

		leftVals = append(leftVals, left)
		rightVals = append(rightVals, right)
	}

	sort.Sort(sort.Float64Slice(leftVals))
	sort.Sort(sort.Float64Slice(rightVals))

	similarity := calculateSimilarity(leftVals, rightVals)

	fmt.Printf("Total similarity: %f\n", similarity)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func calculateSimilarity(leftVals, rightVals []float64) float64 {
	var similarity = 0.0

	rightOccurrences := make(map[float64]int)
	for _, val := range rightVals {
		rightOccurrences[val] = rightOccurrences[val] + 1
	}

	for _, val := range leftVals {
		if nOccurrences, ok := rightOccurrences[val]; ok {
			similarity += val * float64(nOccurrences)
		}
	}

	return similarity
}
