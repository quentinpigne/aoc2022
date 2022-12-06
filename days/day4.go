package days

import (
	"quentinpigne/aoc2022/utils"
	"strings"
)

// Parts Computation

func ComputeDay4Part1(fileLines []string) int {
	count := 0
	for i := 0; i < len(fileLines); i++ {
		assignment1, assignment2 := GetAssignments(fileLines[i])
		if FullyContains(assignment1, assignment2) {
			count++
		}
	}
	return count
}

func ComputeDay4Part2(fileLines []string) int {
	count := 0
	for i := 0; i < len(fileLines); i++ {
		assignment1, assignment2 := GetAssignments(fileLines[i])
		if Overlap(assignment1, assignment2) {
			count++
		}
	}
	return count
}

// Utilities Functions

func GetAssignments(s string) ([]int, []int) {
	assignments := strings.Split(s, ",")
	assignment1 := utils.Map(strings.Split(assignments[0], "-"), utils.ToInt)
	assignment2 := utils.Map(strings.Split(assignments[1], "-"), utils.ToInt)
	return assignment1, assignment2
}

func FullyContains(a1 []int, a2 []int) bool {
	a1FullyContainsA2 := a1[0] <= a2[0] && a1[1] >= a2[1]
	a2FullyContainsA1 := a2[0] <= a1[0] && a2[1] >= a1[1]
	return a1FullyContainsA2 || a2FullyContainsA1
}

func Overlap(a1 []int, a2 []int) bool {
	return a2[0] <= a1[1] && a2[1] >= a1[0]
}
