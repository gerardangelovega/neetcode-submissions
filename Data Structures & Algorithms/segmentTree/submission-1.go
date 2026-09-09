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
		Left: nil,
		Right: nil,
	}
}

type SegmentTree struct {
	root *Segment
}
func NewSegmentTree(nums []int) *SegmentTree {
	var build func([]int, int, int) *Segment
	build = func(nums []int, start, end int) *Segment {
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
	return &SegmentTree{
		root: build(nums, 0, len(nums) - 1),
	}
}

func (st *SegmentTree) Update(index, val int) {
	var update func(*Segment, int, int)
	update = func(root *Segment, i, v int) {
		if i == root.Start && i == root.End {
			root.Sum = v
			return
		}

		m := (root.Start + root.End) / 2

		if i > m {
			update(root.Right, i, v)
		} else {
			update(root.Left, i, v)
		}
		root.Sum = root.Left.Sum + root.Right.Sum
	}
	update(st.root, index, val)
}

func (st *SegmentTree) Query(start, end int) int {
	var query func(*Segment, int, int) int
	query = func(root *Segment, s, e int) int {
		if s == root.Start && e == root.End {
			return root.Sum	
		}

		m := (root.Start + root.End) / 2

		if s > m {
			return query(root.Right, s, e)
		} else if e <= m {
			return query(root.Left, s, e)
		} else {
			return query(root.Left, s, m) + query(root.Right, m + 1, e)
		}
	}
	return query(st.root, start, end)
}
