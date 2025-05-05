package main

import (
	"container/heap"
	"testing"
)

func TestInitCharacterQueue(t *testing.T) {
	var chars CharacterQueue = []Character{
		{
			char:  "a",
			count: 5,
		}, {
			char:  "b",
			count: 4,
		}, {
			char:  "c",
			count: 3,
		}, {
			char:  "d",
			count: 2,
		}, {
			char:  "e",
			count: 1,
		},
	}

	heap.Init(&chars)
	if chars[0].char != "e" {
		t.Errorf("heap Init failed, expected root e, got %s", chars[0].char)
	}

	heap.Push(&chars, Character{char: "f", count: 0})
	if chars[0].char != "f" {
		t.Errorf("heap.Push failed, expected root f, got %s", chars[0].char)
	}

	expectedOrder := []string{"f", "e", "d", "c", "b", "a"}
	for i := range chars.Len() {
		char := heap.Pop(&chars).(Character)
		if char.char != expectedOrder[i] {
			t.Errorf("heap.Pop failed, unexpected order, next value expected %s but got %s", expectedOrder[i], char.char)
		}
	}
}
