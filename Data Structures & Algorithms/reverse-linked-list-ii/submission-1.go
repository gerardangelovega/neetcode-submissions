/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	start, startPrev := head, dummy
	for range left-1 {
		startPrev = start
		start = start.Next
	}

	end, endNext := head, head.Next
	for range right-1 {
		end = endNext
		endNext = end.Next
	}

	startPrev.Next = nil
	end.Next = nil

	reversed := reverse(start)

	startPrev.Next = reversed
	start.Next = endNext

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var root *ListNode

	for head != nil {
		temp := head
		head = head.Next

		temp.Next = root
		root = temp
	}

	return root
}