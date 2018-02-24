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

type stack struct {
	top    *node
	length int
}

type node struct {
	next  *node
	value *TreeNode
}

func (s *stack) push(value *TreeNode) {
	s.top = &(node{s.top, value})
	s.length++
}

func (s *stack) pop() *TreeNode {
	if s.length < 1 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.length--

	return value
}

func preorderTraversal(root *TreeNode) []int {
	output := make([]int, 0)

	todo := stack{nil, 0}
	todo.push(root)

	for todo.length > 0 {
		node := todo.pop()
		if node != nil {
			if node.Left == nil && node.Right == nil {
				output = append(output, node.Val)
			} else {
				todo.push(node.Right)
				todo.push(node.Left)
				todo.push(&TreeNode{node.Val, nil, nil})

			}
		}
	}

	return output
}

func main() {
	n1 := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil}
	fmt.Println(preorderTraversal(n1))
}
