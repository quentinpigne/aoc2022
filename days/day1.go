package days

import (
	"quentinpigne/aoc2022/utils"
	"sort"
	"strconv"
)

// Parts Computation

func ComputeDay1Part1(fileLines []string) int {
	weights := GroupAndSumWeights(fileLines)
	return utils.Max(weights)
}

func ComputeDay1Part2(fileLines []string) int {
	weights := GroupAndSumWeights(fileLines)
	sortedWeights := SortWeights(weights)
	return utils.Sum(sortedWeights[0:3])
}

// Utilities Functions

func GroupAndSumWeights(fileLines []string) []int {
	var weights []int
	i := 0
	weights = append(weights, 0)
	for _, value := range fileLines {
		if value != "" {
			int1, _ := strconv.Atoi(value)
			weights[i] += int1
		} else {
			i++
			weights = append(weights, 0)
		}
	}
	return weights
}

func SortWeights(weights []int) []int {
	sortedWeights := weights[:]
	sort.Sort(sort.Reverse(sort.IntSlice(sortedWeights)))
	return sortedWeights
}
