package main

import (
	"fmt"
	"quentinpigne/aoc2022/days"
	"quentinpigne/aoc2022/utils"
)

func main() {
	fileLines := utils.ReadFileLines("inputs/day5.txt")
	fmt.Println(days.ComputeDay5Part2(fileLines))
}
