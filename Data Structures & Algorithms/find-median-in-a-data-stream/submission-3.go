type Heap struct {
	items []int
	cmp func(int, int) bool
}
func NewHeap(cmp func(int, int) bool) Heap {
	return Heap{
		items: make([]int, 1),
		cmp: cmp,
	}
}
func (h *Heap) Push(val int) {
	h.items = append(h.items, val)
	if len(h.items) == 2 {
		return
	}
	h.SiftUp(len(h.items)-1)
}
func (h *Heap) Pop() int {
	if len(h.items) == 1 {
		return -1
	}
	if len(h.items) == 2 {
		rv := h.items[len(h.items)-1]
		h.items = h.items[:len(h.items)-1]
		return rv
	}
	rv := h.items[1] 
	h.items[1] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.SiftDown(1)
	return rv
}
func (h *Heap) Top() int {
	if len(h.items) == 1 {
		return -1
	}
	return h.items[1]
}
func (h *Heap) SiftUp(i int) {
	for i > 1 && h.cmp(h.items[i], h.items[i >> 1]) {
		h.items[i], h.items[i >> 1] = h.items[i >> 1], h.items[i]
		i = i >> 1
	}
}
func (h *Heap) SiftDown(i int) {
	n := len(h.items)
	for (i << 1) < n {
		l := (i << 1)
		r := (i << 1) | 1
		s :=  i

		if l < n && h.cmp(h.items[l], h.items[s]) {
			s = l
		}
		if r < n && h.cmp(h.items[r], h.items[s]) {
			s = r
		}
		if s == i {
			break
		}
		h.items[i], h.items[s] = h.items[s], h.items[i]
		i = s
	}
}

type MedianFinder struct {
    small Heap // max heap for the small half of the numbers
	large Heap // min heap for the large half of the numbers
}
func Constructor() MedianFinder {
	return MedianFinder{
		small: NewHeap(func(a, b int) bool { return a > b }),
		large: NewHeap(func(a, b int) bool { return a < b }),
	}
}
func (this *MedianFinder) AddNum(num int)  {
    this.small.Push(num)

	if len(this.small.items) > 1 && len(this.large.items) > 1 && this.small.Top() > this.large.Top() { 
		small := this.small.Pop()
		large := this.large.Pop()
		this.small.Push(large)
		this.large.Push(small)
	}

	if len(this.small.items)-len(this.large.items) == 2 {
		this.large.Push(this.small.Pop())
	} else if len(this.small.items)-len(this.large.items) == -2 {
		this.small.Push(this.large.Pop())	
	}
}
func (this *MedianFinder) FindMedian() float64 {
    if len(this.small.items) > len(this.large.items) {
		return float64(this.small.Top())
	} else if len(this.small.items) < len(this.large.items) {
		return float64(this.large.Top())
	} else {
		return (float64(this.small.Top()) + float64(this.large.Top())) / 2.0
	}
}
