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

	letterFreqs, err := countLetterFrequencies(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for k, v := range letterFreqs {
		fmt.Printf("%s occurs %d times\n", string(k), v)
	}

	os.Exit(0)
}
