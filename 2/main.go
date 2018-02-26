package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ptr1, ptr2 := l1, l2
	carry := 0

	var result *ListNode
	var lastNode *ListNode

	for ptr1 != nil || ptr2 != nil || carry != 0 {
		sum := carry
		if ptr1 != nil {
			sum += ptr1.Val
			ptr1 = ptr1.Next
		}
		if ptr2 != nil {
			sum += ptr2.Val
			ptr2 = ptr2.Next
		}

		num := sum % 10
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		node := ListNode{num, nil}

		if result == nil {
			result = &node
		} else {
			lastNode.Next = &node
		}

		lastNode = &node
	}

	return result
}

func main() {

}
