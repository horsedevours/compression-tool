package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 3 {
		fmt.Println("Invalid command")
		fmt.Println("Usage: ./compression-tool <compress|decompress> <inputFile> <outputFile>")
		os.Exit(1)
	}

	operation := os.Args[1]
	input := os.Args[2]
	output := os.Args[3]

	if input == output {
		fmt.Println("Input and output files must be different")
		os.Exit(1)
	}

	switch operation {
	case "compress":
		compressFile(input, output)
	case "decompress":
		decompressFile(input, output)
	default:
		fmt.Printf("Invalid command: %s", operation)
	}

	os.Exit(0)
}
