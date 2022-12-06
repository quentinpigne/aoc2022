package days

import (
	"quentinpigne/aoc2022/utils"
	"regexp"
)

var re = regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

// Parts Computation

func ComputeDay5Part1(fileLines []string) string {
	stacks := initializeStacks(fileLines[:8])

	moves := fileLines[10:]
	for _, move := range moves {
		applyMove(stacks, move)
	}

	result := ""
	for i := 0; i < 9; i++ {
		result += string(stacks[i].Pop())
	}

	return result
}

func ComputeDay5Part2(fileLines []string) string {
	stacks := initializeStacks(fileLines[:8])

	moves := fileLines[10:]
	for _, move := range moves {
		applyMoveOrdered(stacks, move)
	}

	result := ""
	for i := 0; i < 9; i++ {
		result += string(stacks[i].Pop())
	}

	return result
}

// Utilities Functions

func initializeStacks(initialState []string) []*utils.Stack {
	stacks := make([]*utils.Stack, 9)
	for i := 0; i < len(stacks); i++ {
		stacks[i] = &utils.Stack{}
	}
	for i := len(initialState) - 1; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			if initialState[i][1+(j*4)] != 32 {
				stacks[j].Push(initialState[i][1+(j*4)])
			}
		}
	}
	return stacks
}

func applyMove(stacks []*utils.Stack, move string) {
	parts := re.FindStringSubmatch(move)
	for i := 0; i < utils.ToInt(parts[1]); i++ {
		value := stacks[utils.ToInt(parts[2])-1].Pop()
		stacks[utils.ToInt(parts[3])-1].Push(value)
	}
}

func applyMoveOrdered(stacks []*utils.Stack, move string) {
	parts := re.FindStringSubmatch(move)
	toMove := ""
	for i := 0; i < utils.ToInt(parts[1]); i++ {
		toMove += string(stacks[utils.ToInt(parts[2])-1].Pop())
	}
	for i := len(toMove) - 1; i >= 0; i-- {
		stacks[utils.ToInt(parts[3])-1].Push(toMove[i])
	}
}
