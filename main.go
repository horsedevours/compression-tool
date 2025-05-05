package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File argument required")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Could not open file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	letterFreqs, err := countLetterFrequencies(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for k, v := range letterFreqs {
		fmt.Printf("%s occurs %d times\n", string(k), v)
	}

	os.Exit(0)
}
