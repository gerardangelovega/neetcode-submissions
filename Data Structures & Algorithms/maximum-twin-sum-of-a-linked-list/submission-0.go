/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    fast, slow1, slow2 := head, head, head
	res := math.MinInt
	n := 0

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow1 = slow1.Next
		n++
	}

	twin1 := make([]int, n)
	twin2 := make([]int, n)

	i := 0
	for slow1 != nil {
		twin1[i], twin2[i] = slow1.Val, slow2.Val
		slow1, slow2 = slow1.Next, slow2.Next
		i++
	}

	for l, r := 0, n-1; l < n && r >= 0; l, r = l+1, r-1 {
		res = max(res, twin1[l] + twin2[r])
	}

	return res
}