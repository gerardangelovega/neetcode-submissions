/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode {
		Val: -1,
		Next: head,
	}
	from, tail := &dummy.Next, dummy.Next

	count := 0
	for *from != nil && tail != nil {
		if count < k-1 {
			tail = tail.Next
			count++
			continue
		}
		curr := *from
		next := tail.Next 

		tail.Next = nil
		reversed := reverse(nil, curr)

		*from = reversed
		curr.Next = next
		from = &curr.Next
		tail = next

		count = 0
	}

	return dummy.Next
}

func reverse(prev, curr *ListNode) *ListNode {
	if curr == nil {
		return prev
	}
	next := curr.Next
	curr.Next = prev
	return reverse(curr, next)
}
