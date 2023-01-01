package days

import (
	"quentinpigne/aoc2022/utils"
	"strings"
)

// Parts Computation

func ComputeDay8Part1(fileLines []string) int {
	matrix := makeMatrix(fileLines)
	visibleTreesCount := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if isTreeVisible(i, j, matrix) {
				visibleTreesCount++
			}
		}
	}
	return visibleTreesCount
}

func ComputeDay8Part2(fileLines []string) int {
	matrix := makeMatrix(fileLines)
	var scenicScores []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			scenicScores = append(scenicScores, computeScenicScore(i, j, matrix))
		}
	}
	return utils.Max(scenicScores)
}

// Utilities Functions

func makeMatrix(lines []string) [][]string {
	return utils.Map(lines, func(line string) []string {
		return strings.Split(line, "")
	})
}

func isTreeVisible(line int, column int, matrix [][]string) bool {
	return isVisibleTop(line, column, matrix) || isVisibleRight(line, column, matrix) || isVisibleBottom(line, column, matrix) || isVisibleLeft(line, column, matrix)
}

func isVisibleTop(line int, column int, matrix [][]string) bool {
	visible := true
	for i := line - 1; i >= 0; i-- {
		if matrix[i][column] >= matrix[line][column] {
			visible = false
			break
		}
	}
	return visible
}

func isVisibleRight(line int, column int, matrix [][]string) bool {
	visible := true
	for i := column + 1; i < len(matrix[line]); i++ {
		if matrix[line][i] >= matrix[line][column] {
			visible = false
			break
		}
	}
	return visible
}

func isVisibleBottom(line int, column int, matrix [][]string) bool {
	visible := true
	for i := line + 1; i < len(matrix); i++ {
		if matrix[i][column] >= matrix[line][column] {
			visible = false
			break
		}
	}
	return visible
}

func isVisibleLeft(line int, column int, matrix [][]string) bool {
	visible := true
	for i := column - 1; i >= 0; i-- {
		if matrix[line][i] >= matrix[line][column] {
			visible = false
			break
		}
	}
	return visible
}

func computeScenicScore(line int, column int, matrix [][]string) int {
	return lookingTop(line, column, matrix) * lookingRight(line, column, matrix) * lookingBottom(line, column, matrix) * lookingLeft(line, column, matrix)
}

func lookingTop(line int, column int, matrix [][]string) int {
	count := 0
	for i := line - 1; i >= 0; i-- {
		count++
		if matrix[i][column] >= matrix[line][column] {
			break
		}
	}
	return count
}

func lookingRight(line int, column int, matrix [][]string) int {
	count := 0
	for i := column + 1; i < len(matrix[line]); i++ {
		count++
		if matrix[line][i] >= matrix[line][column] {
			break
		}
	}
	return count
}

func lookingBottom(line int, column int, matrix [][]string) int {
	count := 0
	for i := line + 1; i < len(matrix); i++ {
		count++
		if matrix[i][column] >= matrix[line][column] {
			break
		}
	}
	return count
}

func lookingLeft(line int, column int, matrix [][]string) int {
	count := 0
	for i := column - 1; i >= 0; i-- {
		count++
		if matrix[line][i] >= matrix[line][column] {
			break
		}
	}
	return count
}
