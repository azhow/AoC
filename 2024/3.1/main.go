package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	// Read the file
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	input := string(buf)

	programOutput := evaluateProgram(input)

	fmt.Printf("Result: %d\n", programOutput)
}

func evaluateProgram(input string) int {
	programResult := 0

	var validInstruction = regexp.MustCompile(`mul\(([0-9]|[0-9][0-9]|[0-9][0-9][0-9]),([0-9]|[0-9][0-9]|[0-9][0-9][0-9])\)`)

	allMatches := validInstruction.FindAllStringSubmatch(input, -1)
	if allMatches == nil {
		fmt.Println("Error finding all matches")
	}

	for _, match := range allMatches {
		arg1, _ := strconv.Atoi(match[1])
		arg2, _ := strconv.Atoi(match[2])

		programResult += arg1 * arg2
	}

	return programResult
}
