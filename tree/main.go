package main

import (
	"datastruct"
	"fmt"
)

func main() {
	tree := datastruct.Tree{}
	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}

	tree.DFS1()
	fmt.Println()
	fmt.Println()

	tree.DFS2()
	fmt.Println()

	tree.BFS()

}
