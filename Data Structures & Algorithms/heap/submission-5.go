func SlicePush(slice *[]int, val int) {
    *slice = append(*slice, val)
}
func SlicePop(slice *[]int) int {
    n := len(*slice)
    rv := (*slice)[n-1]
    *slice = (*slice)[:n-1]
    return rv
}
func SliceSwap(slice *[]int, a, b int) {
    (*slice)[a], (*slice)[b] = (*slice)[b], (*slice)[a]
}

type MinHeap struct {
    items []int
}
func NewMinHeap() *MinHeap {
    return &MinHeap{
        items: make([]int, 0),
    }
}
func (mh *MinHeap) Push(val int) {
    SlicePush(&mh.items, val)

    if len(mh.items) == 1 {
        return
    }

    mh.SiftUp(len(mh.items)-1);
}
func (mh *MinHeap) Pop() int {
    fmt.Println("Pop")
    n := len(mh.items)
    if n == 0 {
        return -1
    }
    if n == 1 {
        return SlicePop(&mh.items)
    }

    rv := mh.items[0]
    mh.items[0] = SlicePop(&mh.items)
    mh.SiftDown(0)

    return rv
}
func (mh *MinHeap) Top() int {
    if len(mh.items) == 0 {
        return -1
    }
    return mh.items[0]
}
func (mh *MinHeap) Heapify(nums []int) {
    mh.items = make([]int, len(nums))
    copy(mh.items, nums)

    curr := (len(mh.items) - 2) / 2

    for curr > -1 {
        mh.SiftDown(curr)
        curr = curr - 1
    }
}
func (mh *MinHeap) SiftUp(i int) {
    for i > 0 && mh.items[i] < mh.items[mh.Parent(i)] {
        SliceSwap(&mh.items, i, mh.Parent(i))
        i = mh.Parent(i)
    }
}
func (mh *MinHeap) SiftDown(i int) {
    n := len(mh.items)

    for mh.Left(i) < n {
        left := mh.Left(i)
        right := mh.Right(i)
        smallest := i

        if left < n && mh.items[left] < mh.items[smallest] {
            smallest = left
        }
        if right < n && mh.items[right] < mh.items[smallest] {
            smallest = right
        }
        if smallest == i {
            break
        }

        SliceSwap(&mh.items, i, smallest)
        i = smallest
    }
}
func (mh *MinHeap) Left(i int) int {
    return 2 * i + 1
}
func (mh *MinHeap) Right(i int) int {
    return 2 * i + 2
}
func (mh *MinHeap) Parent(i int) int {
    return (i - 1) / 2
}
