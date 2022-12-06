package days

import "quentinpigne/aoc2022/utils"

// Parts Computation

func ComputeDay6Part1(fileLines []string) int {
	return firstMarkerAfterIndex(fileLines[0], 4)
}

func ComputeDay6Part2(fileLines []string) int {
	return firstMarkerAfterIndex(fileLines[0], 14)
}

// Utilities Functions

func firstMarkerAfterIndex(datastream string, markerSize int) int {
	lowerIndex := 0
	upperIndex := markerSize
	for utils.ContainsSameCharacters(datastream[lowerIndex:upperIndex]) {
		lowerIndex++
		upperIndex++
	}
	return upperIndex
}
