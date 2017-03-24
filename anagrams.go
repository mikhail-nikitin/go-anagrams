package go_anagrams

import "unicode"

func IsAnagram(first, second string) (result bool) {
	firstRunes := []rune(first)
	secondRunes := []rune(second)

	if len(firstRunes) != len(secondRunes) {
		return
	}

	firstRuneSet := map[rune]int{}
	secondRuneSet := map[rune]int{}

	for i := 0; i < len(firstRunes); i++ {
		firstRune := unicode.ToUpper(firstRunes[i])
		secondRune := unicode.ToUpper(secondRunes[i])
		firstRuneSet[firstRune]++
		secondRuneSet[secondRune]++
	}

	result = true
	for r, firstRuneCount := range firstRuneSet {
		secondRuneCount, ok := secondRuneSet[r]
		if !ok || firstRuneCount != secondRuneCount {
			result = false
			break
		}
	}
	return
}