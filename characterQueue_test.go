package main

import (
	"container/heap"
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
			q: CharacterQueue{
				{
					val: Character{
						char:  "a",
						count: 3,
					},
				}, {
					val: Character{
						char:  "b",
						count: 2,
					},
				},
			},
			expected: CharacterTree{
				val: Character{
					count: 5,
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
		})
	}
}
