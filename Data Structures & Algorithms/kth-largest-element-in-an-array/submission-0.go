type MaxHeap struct {
	heap []int
}
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: []int{0},
	}
}
func (mh *MaxHeap) Push(val int) {
	if mh.Length() == 0 {
		return
	}
	if mh.Length() == 1 {
		PushArr(&mh.heap, val)
		return
	}

	PushArr(&mh.heap, val)
	mh.siftUp(mh.Length()-1)
}
func (mh *MaxHeap) Pop() int {
	if mh.Length() < 2 {
		return -1
	}
	if mh.Length() == 2 {
		return PopArr(&mh.heap)
	}

	res := mh.heap[1]
	mh.heap[1] = PopArr(&mh.heap)
	mh.siftDown(1)

	return res
}
func (mh *MaxHeap) Top() int {
	if mh.Length() < 2 {
		return -1
	}
	return mh.heap[1]
}
func (mh *MaxHeap) Length() int {
	return len(mh.heap)
}
func (mh *MaxHeap) Heapify(arr []int) {
	mh.heap = make([]int, len(arr)+1)
	copy(mh.heap[1:], arr)

	curr := (len(mh.heap) - 1) / 2

	for curr > 0 {
		i := curr
		mh.siftDown(i)
		curr--
	}
}
func (mh *MaxHeap) siftUp(i int) {
	for i > 1 && mh.heap[i] > mh.heap[mh.parent(i)] {
		SwapArr(&mh.heap, i, mh.parent(i))
		i = mh.parent(i)
	}
}
func (mh *MaxHeap) siftDown(i int) {
	for mh.left(i) < mh.Length() {
		l := mh.left(i)
		r := mh.right(i)
		largest := i

		if l < mh.Length() && mh.heap[l] > mh.heap[largest] {
			largest = l
		}
		if r < mh.Length() && mh.heap[r] > mh.heap[largest] {
			largest = r
		}
		if largest == i {
			break
		}

		SwapArr(&mh.heap, i, largest)
		i = largest
	}
}
func (mh *MaxHeap) parent(i int) int {
	return i / 2
}
func (mh *MaxHeap) left(i int) int {
	return 2 * i
}
func (mh *MaxHeap) right(i int) int {
	return 2 * i + 1
}

func findKthLargest(nums []int, k int) int {
	heap := NewMaxHeap()
	heap.Heapify(nums)

	curr := 0

	for k > 0 {
		curr = heap.Pop()
		k--
	}

	return curr
}

func PushArr(arr *[]int, val int) {
	*arr = append(*arr, val)
}
func PopArr(arr *[]int) int {
	res := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return res
}
func SwapArr(arr *[]int, ia, ib int) {
	(*arr)[ia], (*arr)[ib] = (*arr)[ib], (*arr)[ia]
}