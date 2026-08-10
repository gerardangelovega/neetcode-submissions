func push(arr *[]int, val int) {
	*arr = append(*arr, val)
}

func pop(arr *[]int) int {
	res := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return res
}

func swap(arr *[]int, index1, index2 int) {
	(*arr)[index1], (*arr)[index2] = (*arr)[index2], (*arr)[index1]
}

type MinHeap struct {
	items []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		items: []int{0},
	}
}

func (mh *MinHeap) Push(val int) {
	push(&mh.items, val)

	if len(mh.items) == 1 {
		return
	}

	curr := len(mh.items) - 1
	for curr > 1 && mh.items[curr] < mh.items[curr/2] {
		swap(&mh.items, curr, curr/2)
		curr = curr / 2
	}
}

func (mh *MinHeap) Pop() int {
	if len(mh.items) <= 1 {
		return -1
	}
	if len(mh.items) == 2 {
		return pop(&mh.items)
	}

	res := mh.items[1]
	mh.items[1] = pop(&mh.items)
	curr := 1
	for 2 * curr < len(mh.items) {
		if 2 * curr + 1 < len(mh.items) && mh.items[2*curr+1] < mh.items[2*curr] && mh.items[curr] > mh.items[2*curr+1] {
			swap(&mh.items, curr, 2 * curr + 1)
			curr = 2 * curr + 1
		} else if mh.items[curr] > mh.items[2*curr] {
			swap(&mh.items, curr, 2 * curr)
			curr = 2 * curr
		} else {
			break
		}
	}

	return res
}

func (mh *MinHeap) Top() int {
	if len(mh.items) == 1 {
		return -1
	}

	return mh.items[1]
}

func (mh *MinHeap) Heapify(nums []int) {
	mh.items = make([]int, len(nums) + 1)
	copy(mh.items[1:], nums)

	curr := (len(mh.items) - 1) / 2
	for curr > 0 {
		i := curr
		for 2 * i < len(mh.items) {
			if 2 * i + 1 < len(mh.items) && mh.items[2*i+1] < mh.items[2*i] && mh.items[i] > mh.items[2*i+1] {
				swap(&mh.items, i, 2 * i + 1)
				i = 2 * i + 1
			} else if mh.items[i] > mh.items[2*i] {
				swap(&mh.items, i, 2 * i)
				i = 2 * i
			} else {
				break
			}
		}
		curr = curr - 1
	}
}