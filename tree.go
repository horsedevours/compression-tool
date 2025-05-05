package main

import "sort"

type Tree struct {
	Value  string
	Weight int
	Parent *Tree
	Left   *Tree
	Right  *Tree
}

func MapToTree(data map[string]int) (Tree, error) {
	treeArray := []Tree{}

	for k, v := range data {
		treeArray = append(treeArray, Tree{
			Value:  k,
			Weight: v,
		})
	}

	sort.Slice(treeArray, func(i, j int) bool {
		return treeArray[i].Weight < treeArray[j].Weight
	})

	return Tree{}, nil
}
