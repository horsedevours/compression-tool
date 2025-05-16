package main

import (
	"fmt"
	"os"

	"github.com/horsedevours/compression-tool/internal/bitops"
)

func decompressFile(input string, output string) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("Could not open file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, err := os.Stat(input)
	if err != nil {
		fmt.Printf("Could not obtain input file info: %v", err)
		os.Exit(1)
	}
	inputSize := fileInfo.Size()

	outFile, err := os.Create(output)
	if err != nil {
		fmt.Printf("Could not create file: %v", err)
		os.Exit(1)
	}
	defer outFile.Close()

	err = bitops.BitWiseRead(file, inputSize, outFile)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
