package datastruct

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

//recurrence
func (t *Tree) DFS1() {
	DFS1(t.Root)
}
func DFS1(node *TreeNode) {
	fmt.Printf("%d ", node.Val)
	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}
}

//stack
func (t *Tree) DFS2() {
	var s []*TreeNode
	s = []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last = s[len(s)-1]
		s = s[:len(s)-1]

		fmt.Printf("%d is last \n", last.Val)

		for i := 0; i < len(last.Childs); i++ {
			s = append(s, last.Childs[i])
		}
	}
}

func (t *Tree) BFS() {
	queue := []*TreeNode{}
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		fmt.Printf("%d", first.Val)
		for i := 0; i < len(first.Childs); i++ {
			queue = append(queue, first.Childs[i])
		}
	}
}

//DFS깊이 BFS너비
//DFS 길찾기에 유리함
