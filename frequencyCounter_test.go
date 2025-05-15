package main

import (
	"strings"
	"testing"
)

func TestCountLetterFrequencies(t *testing.T) {
	cases := []struct {
		name           string
		text           string
		expectedChars  []byte
		expectedCounts []int
	}{
		{
			name:           "simple string",
			text:           "aaabbccaaabbcc",
			expectedChars:  []byte{'a', 'b', 'c'},
			expectedCounts: []int{6, 4, 4},
		}, {
			name:           "lots of space",
			text:           "       x    x   ",
			expectedChars:  []byte{' ', 'x'},
			expectedCounts: []int{14, 2},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			freqs := countLetterFrequencies(strings.NewReader(tc.text))

			assertFreqs(freqs, tc.expectedChars, tc.expectedCounts, t)
		})
	}
}

func assertFreqs(freqs map[byte]int, expectedChars []byte, expectedCounts []int, t *testing.T) {
	t.Helper()
	for i, b := range expectedChars {
		if count, ok := freqs[b]; !ok {
			t.Errorf("expected char missing from frequency counts: %b", b)
		} else if count != expectedCounts[i] {
			t.Errorf("expected char %b count to be %d, got %d", b, expectedCounts[i], count)
		}
	}
}
