type Segment struct {
	Start int
	End int
	Sum int
	Left *Segment
	Right *Segment
}
func NewSegment(start, end, sum int) *Segment {
	return &Segment{
		Start: start,
		End: end,
		Sum: sum,
	}
}

type SegmentTree struct {
	root *Segment	
	nums []int
}

func NewSegmentTree(nums []int) *SegmentTree {
	return &SegmentTree{
		root: build(nums, 0, len(nums)-1),
		nums: nums,
	}
}

func build(nums []int, start, end int) *Segment {
	if start == end {
		return NewSegment(start, end, nums[start])
	}

	mid := (start + end) / 2

	root := NewSegment(start, end, 0)
	root.Left = build(nums, start, mid)
	root.Right = build(nums, mid + 1, end)
	root.Sum = root.Left.Sum + root.Right.Sum

	return root
}

func (st *SegmentTree) Update(index, val int) {
	var update func(*Segment, int, int)
	update = func(root *Segment, index, val int) {
		if index == root.Start && index == root.End {
			root.Sum = val
			return
		}

		mid := (root.Start + root.End) / 2

		if index > mid {
			update(root.Right, index, val)
		} else {
			update(root.Left, index, val)
		}
		root.Sum = root.Right.Sum + root.Left.Sum
	}
	update(st.root, index, val)
}

func (st *SegmentTree) Query(start, end int) int {
	var query func(*Segment, int, int) int
	query =	func(root *Segment, s, e int) int {
		if s == root.Start && e == root.End {
			return root.Sum
		}
	
		mid := (root.Start + root.End) / 2
	
		if s > mid {
			return query(root.Right, s, e)
		} else if e <= mid {
			return query(root.Left, s, e)
		} else {
			return query(root.Left, s, mid) + query(root.Right, mid + 1, e)
		}
	}
	return query(st.root, start, end)
}