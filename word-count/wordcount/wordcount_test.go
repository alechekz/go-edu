package wordcount

import (
	"testing"
)

// init test case struct
type testCase struct {
	words string
	count map[string]int
}

// TestCount runs Count function tests
func TestCount(t *testing.T) {

	//init test cases
	var cases []testCase = []testCase{
		{
			words: "some words",
			count: map[string]int{
				"some":  1,
				"words": 1,
			},
		},
		{
			words: "some some words",
			count: map[string]int{
				"some":  2,
				"words": 1,
			},
		},
		{
			words: "some	words with tab",
			count: map[string]int{
				"some":  1,
				"words": 1,
				"with":  1,
				"tab":   1,
			},
		},
		{
			words: "",
			count: map[string]int{},
		},
	}

	//run tests
	for _, tc := range cases {
		got := Count(tc.words)
		if len(got) != len(tc.count) {
			t.Errorf("for input %q: expected %v words, got %v", tc.words, len(tc.count), len(got))
		}
		for word, expectedCount := range tc.count {
			if got[word] != expectedCount {
				t.Errorf("for input %q: word %q expected %d, got %d",
					tc.words, word, expectedCount, got[word])
			}
		}
	}
}
