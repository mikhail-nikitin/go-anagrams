package go_anagrams

import (
	"unicode"
	"sort"
	"strings"
)

func IsAnagram(first, second string) (result bool) {
	firstRunes := []rune(first)
	secondRunes := []rune(second)

	if len(firstRunes) != len(secondRunes) {
		return false
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

func IsAnagramUsingSorting(first, second string) bool {
	return getStringWithSortedCharacters(first) == getStringWithSortedCharacters(second)
}

func getUpperCaseRuneSetWithLength(runes []rune) (runeSet map[rune]int) {
	runeSet = map[rune]int{}
	for _, r := range runes {
		r = unicode.ToUpper(r)
		runeSet[r]++
	}
	return
}

func getStringWithSortedCharacters(s string) string {
	s = strings.ToUpper(s)
	sortedFirstCharacterStrings := strings.Split(s, "")
	sort.Strings(sortedFirstCharacterStrings)
	return strings.Join(sortedFirstCharacterStrings,"")
}