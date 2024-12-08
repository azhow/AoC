package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
        "strconv"
        "slices"
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

	pageOrderings := map[string][]string{}

	// Iterate through each line for page orderings
	for scanner.Scan() {
		line := scanner.Text() // Read the current line as a string

                if line == "" {
                        break
                }

                split := strings.Split(line, "|")
                left := split[0]
                right := split[1]

                pageOrderings[left] = append(pageOrderings[left], right)
	}

        printUpdates := [][]string{}
        // print updates
	for scanner.Scan() {
		line := scanner.Text() // Read the current line as a string

                printUpdates = append(printUpdates, strings.Split(line, ","))
	}

	middlePageSum := sumMiddlePages(pageOrderings, printUpdates)

	fmt.Printf("Result: %d\n", middlePageSum)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func sumMiddlePages(pageOrderings map[string][]string, printUpdates [][]string) int {
        count := 0

        for _, printUpdate := range printUpdates {
                if isUpdateValid(pageOrderings, printUpdate) {
                        count += getMiddlePage(printUpdate)
                }
        }

        return count
}

func isUpdateValid(pageOrderings map[string][]string, update []string) bool {
        result := true

        valuesToTheLeft := []string{}
        for _, val := range update {
                // Check if values are to the left (happened before)
                for _, page := range pageOrderings[val] {
                        result = result && !slices.Contains(valuesToTheLeft, page)
                }
                valuesToTheLeft = append(valuesToTheLeft, val)
        }

        return result
}

func getMiddlePage(update []string) int {
        middlePoint := len(update) / 2

        middlePage, _ := strconv.Atoi(update[middlePoint])

        return middlePage
}
