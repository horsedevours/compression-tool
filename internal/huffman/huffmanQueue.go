package huffman

import (
	"container/heap"
)

type HuffmanNode struct {
	IsLeaf bool
	Val    byte
	Weight int
	Left   *HuffmanNode
	Right  *HuffmanNode
}

type HuffmanQueue []HuffmanNode

func (q HuffmanQueue) Len() int {
	return len(q)
}

func (q HuffmanQueue) Less(i, j int) bool {
	return q[i].Weight < q[j].Weight
}

func (q HuffmanQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *HuffmanQueue) Push(x any) {
	*q = append(*q, x.(HuffmanNode))
}

func (q *HuffmanQueue) Pop() any {
	old := *q
	n := len(old)
	val := old[n-1]
	*q = old[0 : n-1]
	return val
}

func NewHuffmanQueue(byteCounts map[byte]int) (HuffmanQueue, error) {
	q := HuffmanQueue{}
	for k, v := range byteCounts {
		tree := HuffmanNode{
			IsLeaf: true,
			Val:    k,
			Weight: v,
		}
		q = append(q, tree)
	}
	heap.Init(&q)
	return q, nil
}

func (q *HuffmanQueue) BuildHuffmanTree() (HuffmanNode, error) {
	for q.Len() > 1 {
		t1 := heap.Pop(q).(HuffmanNode)
		t2 := heap.Pop(q).(HuffmanNode)

		newNode := HuffmanNode{
			IsLeaf: false,
			Weight: t1.Weight + t2.Weight,
			Left:   &t1,
			Right:  &t2,
		}

		heap.Push(q, newNode)
	}

	return heap.Pop(q).(HuffmanNode), nil
}

func (hTree *HuffmanNode) BuildCodeMap() map[byte][]bool {
	codeMap := map[byte][]bool{}

	getCodes(hTree, codeMap, []bool{})
	return codeMap
}

func getCodes(hTree *HuffmanNode, codeMap map[byte][]bool, code []bool) {
	if hTree.IsLeaf {
		codeMap[hTree.Val] = code
		return
	}

	left := append([]bool{}, code...)
	left = append(left, false)
	getCodes(hTree.Left, codeMap, left)
	right := append([]bool{}, code...)
	right = append(right, true)
	getCodes(hTree.Right, codeMap, right)
}
