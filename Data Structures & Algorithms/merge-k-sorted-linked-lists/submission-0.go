/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)

	if n == 0 {
		return nil
	}

	dummy := &ListNode{}
	list := dummy
	list.Next = lists[0]

	for i := 1; i < n; i++ {
		list1 := list.Next
		list2 := lists[i]
		list.Next = nil

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

		if list1 != nil {
			list.Next = list1
		} else {
			list.Next = list2
		}

		list = dummy
	}

	return dummy.Next
}
