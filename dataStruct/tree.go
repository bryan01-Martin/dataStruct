package dataStruct

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

func (t *Tree) DFSUsingRecursive() {
	DFS1(t.Root)
}

// DFS Recursive
func DFS1(node *TreeNode) {
	fmt.Printf("%d -> ", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}

}

func (t *Tree) DFSUsingStack() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var pop *TreeNode
		pop, s = s[len(s)-1], s[:len(s)-1]
		fmt.Printf("%d -> ", pop.Val)

		for i := len(pop.Childs) - 1; i >= 0; i-- {
			s = append(s, pop.Childs[i])
		}
	}
}

func (t *Tree) BFSUsingQueue() {
	q := []*TreeNode{}
	q = append(q, t.Root)

	for len(q) > 0 {
		var pop *TreeNode
		pop, q = q[0], q[1:]

		fmt.Printf("%d -> ", pop.Val)

		for i := 0; i < len(pop.Childs); i++ {
			q = append(q, pop.Childs[i])
		}

	}
}
