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
				antiNodeLocations := getAllAntinodes(v[i], v[j])
				for _, l := range antiNodeLocations {
					// Only append the ones that are inbounds
					if l.x >= 0 && l.x < maxX && l.y >= 0 && l.y < maxY {
						dupAntinodes[k] = append(dupAntinodes[k], l)
					}
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

func getAllAntinodes(l1, l2 Location) []Location {
	antinodes := []Location{}

	diffY := l2.y - l1.y
	diffX := l2.x - l1.x

	antinodes = append(antinodes, Location{l1.y - diffY, l1.x - diffX})
	antinodes = append(antinodes, Location{l2.y + diffY, l2.x + diffX})

	return antinodes
}
