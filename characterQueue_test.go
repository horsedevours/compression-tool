package main

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestNewCharacterQueue(t *testing.T) {
	freqMap := map[string]int{
		"a": 5,
		"b": 4,
		"c": 3,
		"d": 2,
		"e": 1,
	}

	chars, _ := NewCharacterQueue(freqMap)

	heap.Init(&chars)
	if chars[0].val.char != "e" {
		t.Errorf("heap Init failed, expected root e, got %s", chars[0].val.char)
	}

	heap.Push(&chars, CharacterTree{
		val: Character{
			char:  "f",
			count: 0,
		},
	})
	if chars[0].val.char != "f" {
		t.Errorf("heap.Push failed, expected root f, got %s", chars[0].val.char)
	}

	expectedOrder := []string{"f", "e", "d", "c", "b", "a"}
	for i := range chars.Len() {
		tree := heap.Pop(&chars).(CharacterTree)
		if tree.val.char != expectedOrder[i] {
			t.Errorf("heap.Pop failed, unexpected order, next value expected %s but got %s", expectedOrder[i], tree.val.char)
		}
	}
}

func TestBuildCharacterTree(t *testing.T) {
	testCases := []struct {
		name     string
		q        CharacterQueue
		expected CharacterTree
	}{
		{
			name: "two trees",
			q:    simpleCharacterQueue(),
			expected: CharacterTree{
				val: Character{
					count: 5,
				},
			},
		}, {
			name: "eight trees",
			q:    biggerCharacterQueue(),
			expected: CharacterTree{
				val: Character{
					count: 306,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tree, err := tc.q.BuildHuffmanTree()
			if err != nil {
				t.Errorf("failed to build Huffman tree")
			}
			if tree.val.count != tc.expected.val.count {
				t.Errorf("expected tree val to be %d, got %d instead", tc.expected.val.count, tree.val.count)
			}
			verifyCodes(tree, "", t)
		})
	}
}

func simpleCharacterQueue() CharacterQueue {
	chars := map[string]int{
		"a": 3,
		"b": 2,
	}
	q, _ := NewCharacterQueue(chars)
	heap.Init(&q)
	return q
}

func biggerCharacterQueue() CharacterQueue {
	chars := map[string]int{
		"C": 32,
		"D": 42,
		"E": 120,
		"K": 7,
		"L": 42,
		"M": 24,
		"U": 37,
		"Z": 2,
	}
	q, _ := NewCharacterQueue(chars)
	heap.Init(&q)
	return q
}

func verifyCodes(tree CharacterTree, code string, t *testing.T) {
	t.Helper()
	if tree.isLeaf {
		fmt.Printf("Leaf val is %s and count is %d; code is %s\n", tree.val.char, tree.val.count, code)
		return
	}

	fmt.Printf("Current tree is %d\n", tree.val.count)
	fmt.Printf("Tree val is %d, sum of children is %d\n", tree.val.count, tree.left.val.count+tree.right.val.count)

	verifyCodes(*tree.left, code+"0", t)
	verifyCodes(*tree.right, code+"1", t)
}
