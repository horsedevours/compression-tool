package main

import (
	"strings"
	"testing"
)

func TestCountLetterFrequencies(t *testing.T) {
	cases := []struct {
		name           string
		text           string
		expectedChars  []string
		expectedCounts []int
	}{
		{
			name:           "simple string",
			text:           "aaabbccaaabbcc",
			expectedChars:  []string{"a", "b", "c"},
			expectedCounts: []int{6, 4, 4},
		}, {
			name:           "lots of space",
			text:           "       x    x   ",
			expectedChars:  []string{" ", "x"},
			expectedCounts: []int{14, 2},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			freqs, _ := countLetterFrequencies(strings.NewReader(tc.text))

			assertFreqs(freqs, tc.expectedChars, tc.expectedCounts, t)
		})
	}
}

func assertFreqs(freqs map[string]int, expectedChars []string, expectedCounts []int, t *testing.T) {
	t.Helper()
	for i, char := range expectedChars {
		if count, ok := freqs[char]; !ok {
			t.Errorf("expected char missing from frequency counts: %s", char)
		} else if count != expectedCounts[i] {
			t.Errorf("expected char %s count to be %d, got %d", char, expectedCounts[i], count)
		}
	}
}
