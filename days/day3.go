package days

import (
	"quentinpigne/aoc2022/utils"
)

// Parts Computation

func ComputeDay3Part1(fileLines []string) int {
	dividedLines := utils.Map(fileLines, utils.CutInTwo)
	commonItemInLines := utils.Map(dividedLines, utils.FindCommonItem)
	priorityByLines := utils.Map(commonItemInLines, ConvertRuneToPriority)
	return utils.Sum(priorityByLines)
}

func ComputeDay3Part2(fileLines []string) int {
	groupedLines := utils.StackBy(fileLines, 3)
	commonItemInLines := utils.Map(groupedLines, utils.FindCommonItem)
	priorityByLines := utils.Map(commonItemInLines, ConvertRuneToPriority)
	return utils.Sum(priorityByLines)
}

// Utilities Functions

func ConvertRuneToPriority(r rune) int {
	if r > 96 {
		return int(r - 96)
	} else {
		return int(r - 38)
	}
}
