/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	curr := root
	carr := 0

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carr
		val := sum % 10
		carr = sum / 10

		curr.Val = val

		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil || l2 != nil {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}
	for l1 != nil {
		sum := l1.Val + carr
		val := sum % 10
		carr = sum / 10

		curr.Val = val

		l1 = l1.Next
		if l1 != nil {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}
	for l2 != nil {
		sum := l2.Val + carr
		val := sum % 10
		carr = sum / 10

		curr.Val = val

		l2 = l2.Next
		if l2 != nil {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}
	if carr != 0 {
		curr.Next = &ListNode{
			Val: carr,
		}
	}

	return root
}