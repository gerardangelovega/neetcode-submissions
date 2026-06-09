/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	list := dummy

    for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			list.Next = list1
			list1 = list1.Next
		} else {
			list.Next = list2
			list2 = list2.Next
		}
		list = list.Next
	}

	list.Next = list1
	if list1 == nil {
		list.Next = list2
	}

	return dummy.Next
}