/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -1,
		Next: nil,
	}
	list := dummy
	head1 := list1
	head2 := list2

    for head1 != nil || head2 != nil {
		if head1 == nil {
			list.Next = head2
			break
		}
		if head2 == nil {
			list.Next = head1
			break
		}
		if head1.Val <= head2.Val {
			temp := head1.Next
			list.Next = head1
			list = list.Next
			head1 = temp
		} else {
			temp := head2.Next
			list.Next = head2
			list = list.Next
			head2 = temp
		}
	}

	return dummy.Next
}
