package main

type Character struct {
	char  string
	count int
}

type CharacterQueue []Character

func (q CharacterQueue) Len() int {
	return len(q)
}

func (q CharacterQueue) Less(i, j int) bool {
	return q[i].count < q[j].count
}

func (q CharacterQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *CharacterQueue) Push(x any) {
	*q = append(*q, x.(Character))
}

func (q *CharacterQueue) Pop() any {
	old := *q
	n := len(old)
	char := old[n-1]
	*q = old[0 : n-1]
	return char
}
