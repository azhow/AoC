package main

import (
	"bufio"
	"fmt"
	"math"
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

	var distance = 0.0
	for idx, left := range leftVals {
		distance += math.Abs(left - rightVals[idx])
	}
	fmt.Printf("Total distance: %f\n", distance)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
