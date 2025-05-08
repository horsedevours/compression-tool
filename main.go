package main

import (
	"container/heap"
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

	q, err := NewCharacterQueue(letterFreqs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	heap.Init(&q)

	hTree, err := q.BuildHuffmanTree()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Root has count %d, children sum to %d\n", hTree.val.count, hTree.left.val.count+hTree.right.val.count)

	codeMap := hTree.BuildCodeMap()
	fmt.Printf("%v", codeMap)

	os.Exit(0)
}
