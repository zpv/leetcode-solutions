package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
Input:
1 -> 2 -> 3 -> 4 -> 5 -> 6
k = 3

Output:
3 -> 2 -> 1 -> 6 -> 5 -> 4

`hi` traverses to (k+1)th element.

Let `prev` be the element the next node points to.
Save lo.Next` and then let `lo.Next` be prev.
Set new `lo` to be original `lo.Next`
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	lo := head
	hi := head.Next

	var ret *ListNode = nil
	var tail *ListNode = nil

	count := 1

	for hi != nil {
		for count%k != 0 {
			if hi == nil {
				if ret == nil {
					return head
				} else {
					return ret
				}
			}
			hi = hi.Next
			count++
		}

		first := lo
		prev := hi
		for count != 0 {
			temp := lo.Next
			lo.Next = prev
			prev = lo
			lo = temp
			count--
		}

		if ret == nil {
			ret = prev
		} else {
			tail.Next = prev
		}
		tail = first
		if hi != nil {
			hi = hi.Next
		}
		count++
	}

	return ret
}

func main() {

}
