package main

import (
	"fmt"
	"quentinpigne/aoc2022/days"
	"quentinpigne/aoc2022/utils"
)

func main() {
	fileLines := utils.ReadFileLines("inputs/day8.txt")
	fmt.Println(days.ComputeDay8Part2(fileLines))
}
