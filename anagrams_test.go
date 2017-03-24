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

func BenchmarkIsAnagram(b *testing.B) {
	pairs := []struct {
		first, second string
	}{
		{"лесопромышленность", "солепромышленность"},
		{"старорежимность", "нерасторжимость"},
		{"австралопитек", "ватерполистка"},
		{"просветитель", "терпеливость"},
		{"покраснение", "пенсионерка"},
		{"равновесие", "своенравие"},
		{"полковник", "клоповник"},
		{"стационар", "соратница"},
		{"вертикаль", "кильватер"},
		{"апельсин", "спаниель"},
		{"внимание", "Вениамин"},
	}

	b.Run("Map", func(t *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, pair := range pairs {
				IsAnagram(pair.first, pair.second)
			}
		}
	})
}