package main

type Character struct {
	char  string
	count int
}

type CharacterTree struct {
	val   Character
	left  *CharacterTree
	right *CharacterTree
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
			val: Character{
				char:  k,
				count: v,
			},
		}
		charQ = append(charQ, tree)
	}
	return charQ, nil
}
