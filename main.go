package main

import (
	"fmt"
	"quentinpigne/aoc2022/days"
	"quentinpigne/aoc2022/utils"
)

func main() {
	fileLines := utils.ReadFileLines("inputs/day6.txt")
	fmt.Println(days.ComputeDay6Part2(fileLines))
}
