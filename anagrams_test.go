package go_anagrams

import (
	"testing"
)

type IsAnagramTestCase struct {
	message, first, second string
	expected               bool
}

func TestIsAnagram(t *testing.T) {
	testFunction(t, IsAnagram)
}

func TestIsAnagramUsingSorting(t *testing.T) {
	testFunction(t, IsAnagramUsingSorting)
}

func testFunction(t *testing.T, testedFunction func (first, second string) bool) {
	testCases := getTestCases()

	for _, testCase := range testCases {
		actual := testedFunction(testCase.first, testCase.second)
		if actual != testCase.expected {
			t.Errorf("IsAnagram(%q, %q) == %s, expected %s for %q", testCase.first, testCase.second, actual, testCase.expected, testCase.message)
		}
	}
}

func getTestCases() []IsAnagramTestCase {
	return []IsAnagramTestCase {
		{"equal strings", "лесопромышленность", "лесопромышленность", true},
		{"equal except case", "ЛеСоПрОмЫшЛеНнОсТь", "лЕсОпРоМыШлЕнНоСтЬ", true},
		{"the second string has more characters", "апельсин", "спаниель заборный", false},
		{"the second string is the reversed first", "апельсин", "нисьлепа", true},
		{"diffrent repetions", "ааб", "бба", false},
		{"empty strings", "", "", true},
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	benchmarkFunction(b, IsAnagram)
}

func BenchmarkIsAnagramUsingSorting(b *testing.B) {
	benchmarkFunction(b, IsAnagramUsingSorting)
}

func benchmarkFunction(b *testing.B, testedFunction func (first, second string) bool) {
	pairs := getBenchmarkingPairs()

	for i := 0; i < b.N; i++ {
		for _, pair := range pairs {
			testedFunction(pair.first, pair.second)
		}
	}
}

func getBenchmarkingPairs() []struct{ first, second string } {
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
	return pairs
}