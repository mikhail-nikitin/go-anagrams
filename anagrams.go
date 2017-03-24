package go_anagrams

import "unicode"

func IsAnagram(first, second string) (result bool) {
	firstRunes := []rune(first)
	secondRunes := []rune(second)

	if len(firstRunes) != len(secondRunes) {
		return
	}

	firstRuneSet := getUpperCaseRuneSetWithLength(firstRunes)
	secondRuneSet := getUpperCaseRuneSetWithLength(secondRunes)

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

func getUpperCaseRuneSetWithLength(runes []rune) (runeSet map[rune]int) {
	runeSet = map[rune]int{}
	for _, r := range runes {
		r = unicode.ToUpper(r)
		runeSet[r]++
	}
	return
}