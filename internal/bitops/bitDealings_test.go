package bitops

import (
	"bytes"
	"container/heap"
	"strings"
	"testing"

	"github.com/horsedevours/compression-tool/internal/huffman"
)

func TestBitwiseWrite(t *testing.T) {
	input := "I love to smell tacos for fun in the sun."

	letterFreqs := huffman.CountLetterFrequencies(strings.NewReader(input))
	q, _ := huffman.NewHuffmanQueue(letterFreqs)
	heap.Init(&q)
	hTree, _ := q.BuildHuffmanTree()

	writer := bytes.Buffer{}
	BitwiseWrite(strings.NewReader(input), &writer, hTree)

	decompressed := bytes.Buffer{}
	BitWiseRead(bytes.NewReader(writer.Bytes()), int64(writer.Len()), &decompressed)
	if input != decompressed.String() {
		t.Errorf("String should match after compression and decompression\n  input: %s\n output: %s", input, decompressed.String())
	}
}
