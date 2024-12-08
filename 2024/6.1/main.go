package main

import (
	"bufio"
	"fmt"
	"os"
)

type GuardPosition struct {
	initialX    int
	initialY    int
	currentX    int
	currentY    int
	currentDirX int
	currentDirY int
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

	level := []string{}
	// 0 deg is going "up" the matrix
	guardPosition := GuardPosition{0, 0, 0, 0, 0, -1}

	// Iterate through each line for page orderings
	y := 0
	for scanner.Scan() {
		line := scanner.Text() // Read the current line as a string

		for x, val := range line {
			if string(val) == "^" {
				guardPosition.initialX = x
				guardPosition.currentX = x
				guardPosition.initialY = y
				guardPosition.currentY = y
			}
		}

		level = append(level, line)

		y++
	}

	traversedSpotCount := countTraversedSpots(level, guardPosition)

	fmt.Printf("Result: %d\n", traversedSpotCount)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func countTraversedSpots(level []string, gp GuardPosition) int {
	count := 0

	levelMaxY := len(level) - 1
	levelMaxX := len(level[0]) - 1

	didGuardLeaveLevel := false
	for !didGuardLeaveLevel {
		// Check if move possible
		nextPosX := gp.currentX + gp.currentDirX
		nextPosY := gp.currentY + gp.currentDirY

		didGuardLeaveLevel = (nextPosX > levelMaxX || nextPosX < 0) ||
			(nextPosY > levelMaxY || nextPosY < 0)

		// Mark already traversed
		if string(level[gp.currentY][gp.currentX]) != "X" {
			level[gp.currentY] = level[gp.currentY][:gp.currentX] + "X" + level[gp.currentY][gp.currentX+1:]
			count++
		}

		// Move
		if !didGuardLeaveLevel {
			// Check if can continue moving in the same direction
			if string(level[nextPosY][nextPosX]) == "#" {
				// Change direction
				if gp.currentDirY == -1 {
					gp.currentDirY = 0
					gp.currentDirX = 1
				} else if gp.currentDirY == 1 {
					gp.currentDirY = 0
					gp.currentDirX = -1
				} else if gp.currentDirX == 1 {
					gp.currentDirY = 1
					gp.currentDirX = 0
				} else if gp.currentDirX == -1 {
					gp.currentDirY = -1
					gp.currentDirX = 0
				}
				// Only move if no collision
			} else {
				// Actually move finally
				gp.currentY = nextPosY
				gp.currentX = nextPosX
			}
		}
	}

	return count
}

func printLevel(level []string) {
	for _, line := range level {
		fmt.Println(line)
	}
	fmt.Println("=======================================================================================")
}
