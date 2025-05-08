package main

import (
	"container/heap"
	"fmt"
)

type Character struct {
	char  string
	count int
}

type CharacterTree struct {
	isLeaf bool
	val    Character
	left   *CharacterTree
	right  *CharacterTree
}

type CharacterQueue []CharacterTree

func (q CharacterQueue) Len() int {
	return len(q)
}

func (q CharacterQueue) Less(i, j int) bool {
	return q[i].val.count < q[j].val.count
}

func (q CharacterQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *CharacterQueue) Push(x any) {
	*q = append(*q, x.(CharacterTree))
}

func (q *CharacterQueue) Pop() any {
	old := *q
	n := len(old)
	char := old[n-1]
	*q = old[0 : n-1]
	return char
}

func NewCharacterQueue(charCounts map[string]int) (CharacterQueue, error) {
	charQ := CharacterQueue{}
	for k, v := range charCounts {
		tree := CharacterTree{
			isLeaf: true,
			val: Character{
				char:  k,
				count: v,
			},
		}
		charQ = append(charQ, tree)
	}
	return charQ, nil
}

func (q *CharacterQueue) BuildHuffmanTree() (CharacterTree, error) {
	for q.Len() > 1 {
		t1 := heap.Pop(q).(CharacterTree)
		t2 := heap.Pop(q).(CharacterTree)

		newTree := CharacterTree{
			isLeaf: false,
			val: Character{
				count: t1.val.count + t2.val.count,
			},
			left:  &t1,
			right: &t2,
		}

		heap.Push(q, newTree)
	}

	return heap.Pop(q).(CharacterTree), nil
}

func (cTree *CharacterTree) BuildCodeMap() map[string]string {
	codeMap := map[string]string{}

	getCodes(cTree, codeMap, "0")
	return codeMap
}

func getCodes(cTree *CharacterTree, codeMap map[string]string, code string) {
	if cTree.isLeaf {
		fmt.Printf("Leaf val is %s and count is %d; code is %s\n", cTree.val.char, cTree.val.count, code)
		codeMap[cTree.val.char] = code
		return
	}

	fmt.Printf("Current tree is %d\n", cTree.val.count)
	fmt.Printf("Tree val is %d, sum of children is %d\n", cTree.val.count, cTree.left.val.count+cTree.right.val.count)

	getCodes(cTree.left, codeMap, code+"0")
	getCodes(cTree.right, codeMap, code+"1")
}
