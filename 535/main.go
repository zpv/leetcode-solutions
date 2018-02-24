package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) (mDepth int) {
	maxDepthHelper(root, 1, &mDepth)
	return
}

func maxDepthHelper(root *TreeNode, depth int, max *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if depth > *max {
			*max = depth
		}
		return
	}

	maxDepthHelper(root.Left, depth+1, max)
	maxDepthHelper(root.Right, depth+1, max)
}

func main() {
	n7 := &TreeNode{7, nil, nil}
	n15 := &TreeNode{15, nil, nil}
	n20 := &TreeNode{20, n15, n7}
	n9 := &TreeNode{9, nil, nil}
	n3 := &TreeNode{3, n9, n20}

	fmt.Println(maxDepth(n3))
}
