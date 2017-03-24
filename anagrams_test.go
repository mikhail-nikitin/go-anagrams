package go_anagrams

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		message, first, second string
		expected               bool
	}{
		{"equal strings", "лесопромышленность", "лесопромышленность", true},
		{"equal except case", "ЛеСоПрОмЫшЛеНнОсТь", "лЕсОпРоМыШлЕнНоСтЬ", true},
		{"the second string has more characters", "апельсин", "спаниель заборный", false},
		{"the second string is the reversed first", "апельсин", "нисьлепа", true},
		{"diffrent repetions", "ааб", "бба", false},
		{"empty strings", "","", true},
	}

	for _, testCase := range testCases {
		actual := IsAnagram(testCase.first, testCase.second)
		if actual != testCase.expected {
			t.Errorf("IsAnagram(%q, %q) == %s, expected %s for %q", testCase.first, testCase.second, actual, testCase.expected, testCase.message)
		}
	}
}
