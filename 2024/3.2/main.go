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

	var validInstruction = regexp.MustCompile(`(mul\(([0-9]|[0-9][0-9]|[0-9][0-9][0-9]),([0-9]|[0-9][0-9]|[0-9][0-9][0-9])\)|do\(\)|don't\(\))`)

	allMatches := validInstruction.FindAllStringSubmatch(input, -1)
	if allMatches == nil {
		fmt.Println("Error finding all matches")
	}

	conditionalState := true
	for _, match := range allMatches {
		command := match[0]
		if command == "don't()" {
			conditionalState = false
		} else if command == "do()" {
			conditionalState = true
		} else if conditionalState {
			arg1, _ := strconv.Atoi(match[2])
			arg2, _ := strconv.Atoi(match[3])
			programResult += arg1 * arg2
		}
	}

	return programResult
}
