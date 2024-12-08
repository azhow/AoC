package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

	pageOrderings := map[int][]int{}

	// Iterate through each line for page orderings
	for scanner.Scan() {
		line := scanner.Text() // Read the current line as a string

		if line == "" {
			break
		}

		split := strings.Split(line, "|")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		pageOrderings[left] = append(pageOrderings[left], right)
	}

	printUpdates := [][]int{}
	// print updates
	for scanner.Scan() {
		line := scanner.Text() // Read the current line as a string

		values := []int{}
		for _, val := range strings.Split(line, ",") {
			convertedVal, _ := strconv.Atoi(val)
			values = append(values, convertedVal)
		}

		printUpdates = append(printUpdates, values)
	}

	middlePageSum := sumMiddlePages(pageOrderings, printUpdates)

	fmt.Printf("Result: %d\n", middlePageSum)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func sumMiddlePages(pageOrderings map[int][]int, printUpdates [][]int) int {
	count := 0

	for _, printUpdate := range printUpdates {
		if !isUpdateValid(pageOrderings, printUpdate) {
			fixedUpdate := fixPrint(pageOrderings, printUpdate)
			count += getMiddlePage(fixedUpdate)
		}
	}

	return count
}

func isUpdateValid(pageOrderings map[int][]int, update []int) bool {
	result := true

	valuesToTheLeft := []int{}
	for _, val := range update {
		// Check if values are to the left (happened before)
		for _, page := range pageOrderings[val] {
			result = result && !slices.Contains(valuesToTheLeft, page)
		}
		valuesToTheLeft = append(valuesToTheLeft, val)
	}

	return result
}

func getMiddlePage(update []int) int {
	middlePoint := len(update) / 2

	middlePage := update[middlePoint]

	return middlePage
}

func fixPrint(pageOrderings map[int][]int, update []int) []int {
	newOrder := []int{}

	for idxToInsert, val := range update {
		// Check if values are to the left (happened before)
		for _, page := range pageOrderings[val] {
			idxFound := slices.Index(newOrder, page)

			if idxFound != -1 {
				idxToInsert = int(math.Min(float64(idxToInsert), float64(idxFound)))
			}
		}

		// If none is found, then this is just append
		newOrder = slices.Insert(newOrder, idxToInsert, val)
	}

	return newOrder
}
