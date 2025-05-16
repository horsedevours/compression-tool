package huffman

import (
	"bufio"
	"io"
)

func CountLetterFrequencies(file io.Reader) map[byte]int {
	letterFreqs := map[byte]int{}

	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		if _, ok := letterFreqs[b]; !ok {
			letterFreqs[b] = 1
		} else {
			letterFreqs[b]++
		}
	}

	return letterFreqs
}
