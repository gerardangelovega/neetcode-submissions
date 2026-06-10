/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

	for curr, n := head, 0; curr != nil; curr, n = curr.Next, n + 1 {
		if n & 1 == 1 {
			continue
		}
		reorder(curr)
	}
}

func reorder(node *ListNode) {
	next := node.Next

	if next == nil {
		return
	}
	if next.Next == nil {
		return
	}
	
	for curr := next; curr != nil; curr = curr.Next {
		if curr.Next.Next == nil {
			temp := curr.Next
			curr.Next = nil
			temp.Next = next
			node.Next = temp
			break
		}
	}
}
