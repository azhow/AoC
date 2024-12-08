package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	result   int
	operands []int
}

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

	equations := []Operation{}

	for scanner.Scan() {
		line := scanner.Text() // Read the current line as a string

		op := Operation{}

		colonSplit := strings.Split(line, ":")
		op.result, _ = strconv.Atoi(colonSplit[0])

		for _, v := range strings.Split(line, " ") {
			operand, _ := strconv.Atoi(v)
			op.operands = append(op.operands, operand)
		}

		equations = append(equations, op)
	}

	sumTotalCalibrationResult := getTotalCalibrationResult(equations)

	fmt.Printf("Result: %d\n", sumTotalCalibrationResult)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func getTotalCalibrationResult(equations []Operation) int {
	result := 0

	for _, e := range equations {
		previousResults := []int{e.operands[0]}
		e.operands = e.operands[1:]
		if isEquationValid(e, previousResults) {
			result += e.result
		}
	}

	return result
}

func isEquationValid(e Operation, previousResults []int) bool {
	if len(e.operands) == 0 {
		resultReached := false
		for _, possibleResult := range previousResults {
			resultReached = resultReached || (possibleResult == e.result)
		}
		return resultReached
	}

	newResults := []int{}
	for _, r := range previousResults {
		addResult := r + e.operands[0]
		mulResult := r * e.operands[0]
		concatStr := strconv.Itoa(r) + strconv.Itoa(e.operands[0])
		concatResult, _ := strconv.Atoi(concatStr)

		newResults = append(newResults, addResult)
		newResults = append(newResults, mulResult)
		newResults = append(newResults, concatResult)
	}
	e.operands = e.operands[1:]

	return isEquationValid(e, newResults)
}
