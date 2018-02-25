package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compareNodes(root.Left, root.Right) && isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(leftTree *TreeNode, rightTree *TreeNode) bool {
	if leftTree == nil || rightTree == nil {
		return leftTree == rightTree
	}
	return compareNodes(leftTree.Left, rightTree.Right) && compareNodes(leftTree.Right, rightTree.Left) && isSymmetricHelper(leftTree.Right, rightTree.Left) && isSymmetricHelper(leftTree.Left, rightTree.Right)
}

func compareNodes(n1, n2 *TreeNode) bool {
	if n1 == nil || n2 == nil {
		return n1 == n2
	}
	return n1.Val == n2.Val
}

func main() {
	n1 := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil}
	fmt.Println(isSymmetric(n1))
}
