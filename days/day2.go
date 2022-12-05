package days

import (
	"quentinpigne/aoc2022/utils"
)

var PointsByShape = map[string]int{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
}

var PointsByHitPart1 = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

var PointsByHitPart2 = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}

// Parts Computation

func ComputeDay2Part1(fileLines []string) int {
	return utils.Sum(ComputePoints(PointsByHitPart1, fileLines))
}

func ComputeDay2Part2(fileLines []string) int {
	return utils.Sum(ComputePoints(PointsByHitPart2, fileLines))
}

// Utilities Functions

func PointsByHitFn(PointsByHit map[string]int) func(line string) int {
	return func(line string) int {
		return PointsByHit[line]
	}
}

func ComputePoints(PointsByHit map[string]int, hits []string) []int {
	return utils.Map(hits, PointsByHitFn(PointsByHit))
}
