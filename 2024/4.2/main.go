package main

import (
	"bufio"
	"fmt"
	"os"
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

	matrix := []string{}
	// Iterate through each line
	for scanner.Scan() {
		row := scanner.Text() // Read the current line as a string

		matrix = append(matrix, row)
	}

	xmasCount := countXmas(matrix)

	fmt.Printf("Result: %d\n", xmasCount)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func countXmas(matrix []string) int {
	var count = 0

	nRows := len(matrix)
	nCols := len(matrix[0])

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			countXmasFromHere2(matrix, i, j, &count)
		}
	}

	return count
}

func countXmasFromHere2(matrix []string, i, j int, count *int) {
	nRows := len(matrix)
	nCols := len(matrix[0])

	xMasStr := "MAS"
	xMasRevStr := "SAM"

	if i < 0 || i+2 >= nRows {
		return
	}

	if j < 0 || j+2 >= nCols {
		return
	}

	// Check if possible matches in current row window
	if (matrix[i][j] == xMasStr[0] || matrix[i][j] == xMasRevStr[0]) &&
		(matrix[i][j+2] == xMasStr[0] || matrix[i][j+2] == xMasRevStr[0]) {
		matchCount := 0

		if matrix[i][j] == xMasStr[0] {
			countXmasFromHere(matrix, i, j, xMasStr, &matchCount, 1, 1)
		} else {
			countXmasFromHere(matrix, i, j, xMasRevStr, &matchCount, 1, 1)
		}

		if matrix[i][j+2] == xMasStr[0] {
			countXmasFromHere(matrix, i, j+2, xMasStr, &matchCount, 1, -1)
		} else {
			countXmasFromHere(matrix, i, j+2, xMasRevStr, &matchCount, 1, -1)
		}

		if matchCount == 2 {
			*count += 1
		}
	}
}

func countXmasFromHere(matrix []string, i, j int, searchStr string, count *int, directionY, directionX int) {
	// Here we want to check only possible matches to the right and down:
	// meaning: (i, j+1), (i+1, j-1), (i+1, j), (i+1, j+1) - maybe check if makes sense the
	// +1 and -1 given the sizes of the matrix

	nRows := len(matrix)
	nCols := len(matrix[0])

	if i < 0 || i >= nRows {
		return
	}

	if j < 0 || j >= nCols {
		return
	}

	if len(searchStr) == 1 && matrix[i][j] == searchStr[0] {
		*count += 1
		return
	}

	if matrix[i][j] == searchStr[0] && directionX == 0 && directionY == 0 {
		countXmasFromHere(matrix, i, j+1, searchStr[1:], count, 0, 1)
		countXmasFromHere(matrix, i+1, j-1, searchStr[1:], count, 1, -1)
		countXmasFromHere(matrix, i+1, j, searchStr[1:], count, 1, 0)
		countXmasFromHere(matrix, i+1, j+1, searchStr[1:], count, 1, 1)
	} else if matrix[i][j] == searchStr[0] {
		countXmasFromHere(matrix, i+directionY, j+directionX, searchStr[1:], count, directionY, directionX)
	}
}
