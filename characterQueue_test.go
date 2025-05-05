package main

import (
	"container/heap"
	"testing"
)

func TestNewCharacterQueue(t *testing.T) {
	// var chars CharacterQueue = []Character{
	// 	{
	// 		char:  "a",
	// 		count: 5,
	// 	}, {
	// 		char:  "b",
	// 		count: 4,
	// 	}, {
	// 		char:  "c",
	// 		count: 3,
	// 	}, {
	// 		char:  "d",
	// 		count: 2,
	// 	}, {
	// 		char:  "e",
	// 		count: 1,
	// 	},
	// }

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
