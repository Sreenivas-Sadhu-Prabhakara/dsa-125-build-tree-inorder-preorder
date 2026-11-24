package main

import (
	"reflect"
	"testing"
)

func preorderOf(n *TreeNode) []int {
	if n == nil {
		return nil
	}
	out := []int{n.Val}
	out = append(out, preorderOf(n.Left)...)
	return append(out, preorderOf(n.Right)...)
}

func TestBuild(t *testing.T) {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	root := Build(pre, in)
	// Round trip: the rebuilt tree must reproduce both traversals.
	if got := preorderOf(root); !reflect.DeepEqual(got, pre) {
		t.Errorf("preorder = %v, want %v", got, pre)
	}
	if got := inorderOf(root); !reflect.DeepEqual(got, in) {
		t.Errorf("inorder = %v, want %v", got, in)
	}

	if Build(nil, nil) != nil {
		t.Error("empty input should build nil")
	}
	single := Build([]int{5}, []int{5})
	if single == nil || single.Val != 5 || single.Left != nil || single.Right != nil {
		t.Error("single-node build wrong")
	}
	// Left-skewed: preorder 3 2 1, inorder 1 2 3.
	skew := Build([]int{3, 2, 1}, []int{1, 2, 3})
	if skew.Val != 3 || skew.Left.Val != 2 || skew.Left.Left.Val != 1 || skew.Right != nil {
		t.Error("left-skewed build wrong")
	}
}
