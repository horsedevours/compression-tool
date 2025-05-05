package main

import (
	"bufio"
	"fmt"
	"io"
)

func countLetterFrequencies(file io.Reader) (map[string]int, error) {
	letterFreqs := map[string]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, b := range line {
			if v, ok := letterFreqs[string(b)]; !ok {
				letterFreqs[string(b)] = 1
			} else {
				letterFreqs[string(b)] = v + 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return map[string]int{}, fmt.Errorf("Scanner error: %w", err)
	}

	return letterFreqs, nil
}
