package main

import (
	"fmt"
	"quentinpigne/aoc2022/days"
	"quentinpigne/aoc2022/utils"
)

func main() {
	fileLines := utils.ReadFileLines("inputs/day3.txt")
	fmt.Println(days.ComputeDay3Part2(fileLines))
}
