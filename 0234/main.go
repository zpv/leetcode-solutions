package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	single, double := head, head
	var prev, prev2 *ListNode
	for {
		double = double.Next.Next

		// single node reverses its previous list
		prev = single
		single = single.Next
		prev.Next = prev2
		prev2 = prev

		if double == nil || double.Next == nil {
			break
		}
	}

	left, right := prev, single
	if double != nil { // adjust needed a number of nodes is even
		right = right.Next
	}

	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}
	return true
}

func main() {

}
