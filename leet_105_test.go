package main

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test001(t *testing.T) {
	var preorder []int = []int{3, 9, 20, 15, 7}
	var inorder []int = []int{9, 3, 15, 20, 7}
	result := buildTree(preorder, inorder)

	if result.Val != 3 {
		t.Error("Wrong Result")
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return partition(0, 0, len(preorder)-1, preorder, inorder)
}

func partition(rootIndex int, start int, end int, preorder []int, inorder []int) *TreeNode {
	if start > end || rootIndex > len(preorder)-1 {
		return nil
	}

	var root *TreeNode = &TreeNode{
		Val:   preorder[rootIndex],
		Left:  nil,
		Right: nil,
	}

	var middle int = 0
	for i := start; i <= end; i++ {
		if inorder[i] == root.Val {
			middle = i
		}
	}

	var countOfChildrenLeft int = middle - start

	root.Left = partition(rootIndex+1, start, middle-1, preorder, inorder)
	root.Right = partition(rootIndex+countOfChildrenLeft+1, middle+1, end, preorder, inorder)
	return root
}
