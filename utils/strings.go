package utils

import (
	"math"
)

func CutInTwo(s string) []string {
	middleIndice := len(s) / 2
	return []string{s[:middleIndice], s[middleIndice:]}
}

func FindCommonItem(data []string) rune {
	arrayMin := [26]int{}
	arrayMaj := [26]int{}
	mainMin := [26]int{}
	mainMaj := [26]int{}

	for i := 0; i < 26; i++ {
		mainMin[i] = math.MaxInt
		mainMaj[i] = math.MaxInt
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] > 96 {
				arrayMin[data[i][j]-'a']++
			} else {
				arrayMaj[data[i][j]-'A']++
			}
		}

		for i := 0; i < 26; i++ {
			if arrayMin[i] < mainMin[i] {
				mainMin[i] = arrayMin[i]
			}
			arrayMin[i] = 0
			if arrayMaj[i] < mainMaj[i] {
				mainMaj[i] = arrayMaj[i]
			}
			arrayMaj[i] = 0
		}
	}

	var result rune
	for i := 0; i < 26; i++ {
		if mainMin[i] > 0 {
			result = rune(i) + 'a'
		}
		if mainMaj[i] > 0 {
			result = rune(i) + 'A'
		}
	}
	return result
}

func ContainsSameCharacters(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}
