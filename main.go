package main

import (
	"fmt"
	"quentinpigne/aoc2022/days"
	"quentinpigne/aoc2022/utils"
)

func main() {
	fileLines := utils.ReadFileLines("inputs/day4.txt")
	fmt.Println(days.ComputeDay4Part2(fileLines))
}
