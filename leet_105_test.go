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

	if result.Left.Val != 9 {
		t.Error("Wrong Result")
	}

	if result.Right.Val != 20 {
		t.Error("Wrong Result")
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return partition(&preorder, inorder)
}

func partition(preorder *[]int, inorder []int) *TreeNode {
	if len(*preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}

	var rootValue = pop(preorder)

	var middle int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootValue {
			middle = i
		}
	}

	var root *TreeNode = &TreeNode{
		Val: rootValue,
	}

	root.Left = partition(preorder, inorder[:middle])
	root.Right = partition(preorder, inorder[middle+1:])
	return root
}

func pop(list *[]int) int {
	rv := (*list)[0]
	*list = (*list)[1:]
	return rv
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	return partition(0, 0, len(preorder) - 1, preorder, inorder)
//}
//
//func partition(rootIndex int, start int, end int, preorder []int, inorder []int) *TreeNode {
//	if start > end || rootIndex > len(preorder)-1 {
//		return nil
//	}
//
//	var root *TreeNode = &TreeNode{
//		Val: preorder[rootIndex],
//		Left: nil,
//		Right: nil,
//	}
//
//	var middle int = 0
//	for i := start; i <= end; i++ {
//		if inorder[i] == root.Val {
//			middle = i
//		}
//	}
//
//	var countOfChildrenLeft int = middle - start
//
//	root.Left = partition(rootIndex+1, start, middle-1, preorder, inorder)
//	root.Right = partition(rootIndex + countOfChildrenLeft + 1, middle + 1, end, preorder, inorder)
//	return root
//}
