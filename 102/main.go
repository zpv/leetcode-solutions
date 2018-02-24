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

type queue struct {
	first  *node
	last   *node
	length int
}

type node struct {
	next  *node
	value interface{}
}

func (s *queue) push(value interface{}) {
	newNode := &(node{nil, value})
	if s.last == nil {
		s.last = newNode
	} else {
		s.last.next = newNode
		s.last = s.last.next
	}
	if s.first == nil {
		s.first = s.last
	}
	s.length++
}

func (s *queue) pop() interface{} {
	if s.length < 1 {
		return nil
	}
	value := s.first.value
	s.first = s.first.next
	s.length--

	if s.length < 1 {
		s.last = nil
	}

	return value
}

func levelOrder(root *TreeNode) [][]int {
	todo := queue{nil, nil, 0}
	depthQueue := queue{nil, nil, 0}

	todo.push(root)
	depthQueue.push(0)

	result := make([][]int, 0)

	for todo.length > 0 {
		node := todo.pop().(*TreeNode)
		depth := depthQueue.pop().(int)
		if node != nil {
			if depth >= len(result) {
				result = append(result, []int{})
			}
			result[depth] = append(result[depth], node.Val)
			todo.push(node.Left)
			todo.push(node.Right)
			depthQueue.push(depth + 1)
			depthQueue.push(depth + 1)
		}
	}
	return result
}

func main() {
	n7 := &TreeNode{7, nil, nil}
	n15 := &TreeNode{15, nil, nil}
	n20 := &TreeNode{20, n15, n7}
	n9 := &TreeNode{9, nil, nil}
	n3 := &TreeNode{3, n9, n20}

	fmt.Println(levelOrder(n3))
}
