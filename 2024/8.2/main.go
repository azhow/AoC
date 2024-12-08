package main

import (
	"bufio"
	"fmt"
	"os"
)

type Location struct {
	y int
	x int
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

	antennasLocations := map[string][]Location{}

	y := 0
	maxX := 0
	for scanner.Scan() {
		line := scanner.Text() // Read the current line as a string
		maxX = len(line)

		for x, v := range line {
			if string(v) != "." {
				antennasLocations[string(v)] = append(antennasLocations[string(v)], Location{y, x})
			}
		}
		y++
	}

	uniqueAntinodeCount := getAllUniqueAntinodesCount(antennasLocations, maxX, y)

	fmt.Printf("Result: %d\n", uniqueAntinodeCount)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func getAllUniqueAntinodesCount(antennasLocations map[string][]Location, maxX, maxY int) int {
	result := 0

	// Get all anti nodes for every pair of antennas
	dupAntinodes := map[string][]Location{}
	for k, v := range antennasLocations {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				antiNodeLocations := getAllAntinodes(v[i], v[j], maxX, maxY)
				for _, l := range antiNodeLocations {
					dupAntinodes[k] = append(dupAntinodes[k], l)
				}
			}
		}
	}

	// Remove duplicates
	allLocations := make(map[Location]bool)
	uniqueAntinodes := map[string][]Location{}
	for k, v := range dupAntinodes {
		for _, l := range v {
			if _, value := allLocations[l]; !value {
				allLocations[l] = true
				uniqueAntinodes[k] = append(uniqueAntinodes[k], l)
			}
		}
	}

	// Count
	for _, v := range uniqueAntinodes {
		result += len(v)
	}

	return result
}

func getAllAntinodes(l1, l2 Location, maxX, maxY int) []Location {
	antinodes := []Location{}

	diffY := l2.y - l1.y
	diffX := l2.x - l1.x

	inRange := true
	mult := 0

	// Calculate the resonant frequencies
	for inRange {
		possibleLocation1 := Location{l1.y - mult * diffY, l1.x - mult * diffX}
		inRange1 := checkInbound(possibleLocation1, maxX, maxY)
		if inRange1 {
			antinodes = append(antinodes, possibleLocation1)
		}

		possibleLocation2 := Location{l2.y + mult * diffY, l2.x + mult * diffX}
		inRange2 := checkInbound(possibleLocation2, maxX, maxY)
		if inRange2 {
			antinodes = append(antinodes, possibleLocation2)
		}

		inRange = inRange1 || inRange2
		mult++
	}

	return antinodes
}

func checkInbound(l Location, maxX, maxY int) bool {
	return l.x >= 0 && l.x < maxX && l.y >= 0 && l.y < maxY
}
