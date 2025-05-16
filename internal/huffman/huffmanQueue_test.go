package huffman

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestNewCharacterQueue(t *testing.T) {
	freqMap := map[byte]int{
		'a': 5,
		'b': 4,
		'c': 3,
		'd': 2,
		'e': 1,
	}

	chars, _ := NewHuffmanQueue(freqMap)

	heap.Init(&chars)
	if chars[0].Val != 'e' {
		t.Errorf("heap Init failed, expected root e, got %b", chars[0].Val)
	}

	heap.Push(&chars, HuffmanNode{
		Val:    'f',
		Weight: 0,
	})
	if chars[0].Val != 'f' {
		t.Errorf("heap.Push failed, expected root f, got %b", chars[0].Val)
	}

	expectedOrder := []byte{'f', 'e', 'd', 'c', 'b', 'a'}
	for i := range chars.Len() {
		tree := heap.Pop(&chars).(HuffmanNode)
		if tree.Val != expectedOrder[i] {
			t.Errorf("heap.Pop failed, unexpected order, next value expected %b but got %b", expectedOrder[i], tree.Val)
		}
	}
}

func TestBuildHuffmanTree(t *testing.T) {
	testCases := []struct {
		name     string
		q        HuffmanQueue
		expected HuffmanNode
	}{
		{
			name: "two trees",
			q:    simpleCharacterQueue(),
			expected: HuffmanNode{
				Weight: 5,
			},
		}, {
			name: "eight trees",
			q:    biggerHuffmanQueue(),
			expected: HuffmanNode{
				Weight: 306,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tree, err := tc.q.BuildHuffmanTree()
			if err != nil {
				t.Errorf("failed to build Huffman tree")
			}
			if tree.Val != tc.expected.Val {
				t.Errorf("expected tree val to be %d, got %d instead", tc.expected.Weight, tree.Weight)
			}
			verifyCodes(tree, "", t)
		})
	}
}

func simpleCharacterQueue() HuffmanQueue {
	chars := map[byte]int{
		'a': 3,
		'b': 2,
	}
	q, _ := NewHuffmanQueue(chars)
	heap.Init(&q)
	return q
}

func biggerHuffmanQueue() HuffmanQueue {
	chars := map[byte]int{
		'C': 32,
		'D': 42,
		'E': 120,
		'K': 7,
		'L': 42,
		'M': 24,
		'U': 37,
		'Z': 2,
	}
	q, _ := NewHuffmanQueue(chars)
	heap.Init(&q)
	return q
}

// Used for manual inspection during development
func verifyCodes(tree HuffmanNode, code string, t *testing.T) {
	t.Helper()
	if tree.IsLeaf {
		fmt.Printf("Leaf val is %b and count is %d; code is %v\n", tree.Val, tree.Weight, code)
		return
	}

	fmt.Printf("Current tree is %d\n", tree.Weight)
	fmt.Printf("Tree val is %d, sum of children is %d\n", tree.Weight, tree.Left.Weight+tree.Right.Weight)

	verifyCodes(*tree.Left, code+"0", t)
	verifyCodes(*tree.Right, code+"1", t)
}
