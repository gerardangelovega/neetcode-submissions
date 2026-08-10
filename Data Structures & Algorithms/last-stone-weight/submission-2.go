type MaxHeap struct {
	heap []int	
}
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: []int{0},
	}
}
func (mh *MaxHeap) Push(val int) {
	// SIMPLY APPEND IF MIN HEAP IS EMPTY
	if mh.Length() <= 1 {
		ArrayPush(&mh.heap, val)
		return
	}

	// APPEND AND SIFT UP IF MIN HEAP IS NOT EMPTY
	ArrayPush(&mh.heap, val)
	mh.siftDown(mh.Length()-1)
}
func (mh *MaxHeap) Pop() int {
	// RETURN -1 IF MIN HEAP IS EMPTY
	if mh.Length() <= 1 {
		return -1
	}
	// RETURN THE FIRST ELEMENT IF MIN HEAP ONLY HAS ONE ITEM
	if mh.Length() == 2 {
		return ArrayPop(&mh.heap)
	}

	// SWAP THE FIRST ELEMENT W/ LAST ELEMENT, SIFT DOWN, AND
	// RETURN THE ORIGINAL FIRST ELEMENT
	res := mh.heap[1]
	mh.heap[1] = ArrayPop(&mh.heap)
	mh.siftDown(1)

	return res
}
func (mh *MaxHeap) Top() int {
	if mh.Length() <= 1 {
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
	for i > 0 && mh.heap[i] > mh.heap[mh.parent(i)] {
		ArraySwap(&mh.heap, i, mh.parent(i))
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

		ArraySwap(&mh.heap, i, largest)
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

func lastStoneWeight(stones []int) int {
	heap := NewMaxHeap()
	heap.Heapify(stones)

	curr := heap.Pop()

	for heap.Length() > 1 {
		if curr == heap.Top() {
			_ = heap.Pop()
		} else if curr < heap.Top() {
			heap.Push(heap.Pop() - curr)
		} else {
			heap.Push(curr - heap.Pop())
		}
		curr = heap.Pop()
	}

	return max(0, curr)
}

func ArrayPush(arr *[]int, val int) {
	*arr = append(*arr, val)	
}
func ArrayPop(arr *[]int) int {
	res := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return res
}
func ArraySwap(arr *[]int, ia, ib int) {
	(*arr)[ia], (*arr)[ib] = (*arr)[ib], (*arr)[ia]
}