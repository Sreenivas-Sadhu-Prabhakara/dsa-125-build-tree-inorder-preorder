package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Build reconstructs the unique binary tree with the given preorder and
// inorder sequences (values distinct). Preorder hands out roots in order;
// inorder says how many nodes belong on each side of that root.
func Build(preorder, inorder []int) *TreeNode {
	at := make(map[int]int, len(inorder))
	for i, v := range inorder {
		at[v] = i
	}
	pre := 0
	var build func(inLo, inHi int) *TreeNode // inorder window [inLo, inHi)
	build = func(inLo, inHi int) *TreeNode {
		if inLo >= inHi {
			return nil
		}
		v := preorder[pre]
		pre++
		n := &TreeNode{Val: v}
		mid := at[v]
		n.Left = build(inLo, mid) // consumes exactly the left-size prefix of preorder
		n.Right = build(mid+1, inHi)
		return n
	}
	return build(0, len(inorder))
}

func inorderOf(n *TreeNode) []int {
	if n == nil {
		return nil
	}
	return append(append(inorderOf(n.Left), n.Val), inorderOf(n.Right)...)
}

func main() {
	root := Build([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
	fmt.Println(inorderOf(root))
}
