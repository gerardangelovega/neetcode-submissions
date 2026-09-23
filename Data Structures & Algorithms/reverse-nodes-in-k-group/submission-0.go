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
		// curr will be the head of the segment, tail will be the tail of the segment
		curr := *from
		
		// temporarily store the ListNode after the tail for reattaching of segments later
		next := tail.Next 

		reversed := reverse(nil, detach(curr, tail))

		// tail becomes the head
		// head becomes the tail
		// attach *from.Next to tail
		// attach curr.Next to next
		*from = reversed
		curr.Next = next
		from = &curr.Next
		tail = next

		count = 0
	}

	return dummy.Next
}

func detach(head *ListNode, tail *ListNode) *ListNode {
	tail.Next = nil
	return head
}

func reverse(prev, curr *ListNode) *ListNode {
	if curr == nil {
		return prev
	}
	next := curr.Next
	curr.Next = prev
	return reverse(curr, next)
}
// SIMULATION for reverse():
// 1>2>3>4>5>nil		start
// 1>nil 2>3>4>5>nil	p:nil	c:1
// 2>1>nil 3>4>5>nil	p:1 	c:2
// 3>2>1>nil 4>5>nil	p:2		c:3
// 4>3>2>1>nil 5>nil	p:3		c:4
// 5>4>3>2>1>nil nil	p:4		c:5
// return				p:5 	c:nil