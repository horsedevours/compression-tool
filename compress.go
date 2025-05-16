package main

import (
	"fmt"
	"io"
	"os"

	"github.com/horsedevours/compression-tool/internal/bitops"
	"github.com/horsedevours/compression-tool/internal/huffman"
)

func compressFile(input string, output string) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("Could not open file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	letterFreqs := huffman.CountLetterFrequencies(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file.Seek(0, io.SeekStart)

	q, err := huffman.NewHuffmanQueue(letterFreqs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hTree, err := q.BuildHuffmanTree()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outFile, err := os.Create(output)
	if err != nil {
		fmt.Printf("unable to create output file %v", err)
		os.Exit(1)
	}
	defer outFile.Close()

	err = bitops.BitwiseWrite(file, outFile, hTree)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
