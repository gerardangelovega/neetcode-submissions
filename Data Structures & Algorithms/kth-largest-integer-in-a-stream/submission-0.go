type KthLargest struct {
	kth int
	heap []int
}
func Constructor(k int, nums []int) KthLargest {
	heap := make([]int, len(nums)+1)
	copy(heap[1:], nums)
	n := len(heap)

	for curr := n/2; curr > 0; curr-- {
		i := curr
		for 2 * i < n {
			l := 2 * i
			r := 2 * i + 1
			largest := i

			if l < n && heap[l] > heap[largest] {
				largest = l
			}
			if r < n && heap[r] > heap[largest] {
				largest = r
			}
			if largest == i {
				break
			}

			heap[i], heap[largest] = heap[largest], heap[i]
			i = largest
		}
	}

	fmt.Println("Finished Heapifying", heap)

	return KthLargest{
		kth: k,
		heap: heap,
	} 
}
func (this *KthLargest) Add(val int) int {
    this.Push(val)

	temp := make([]int, 0, this.kth)

	for i := this.kth; i > 0; i-- {
		temp = append(temp, this.Pop())
	}

	for _, num := range temp {
		this.Push(num)
	}

	return temp[len(temp)-1]
}
func (this *KthLargest) Push(val int) {
	if this.Length() == 0 {
		this.heap = append(this.heap, 0)
		this.heap = append(this.heap, val)
		return
	}
	if this.Length() == 1 {
		this.heap = append(this.heap, val)
		return
	}
	this.heap = append(this.heap, val)
	for i := this.Length() - 1; i > 1 && this.heap[i] > this.heap[this.parent(i)]; i = this.parent(i) {
		this.heap[i], this.heap[this.parent(i)] = this.heap[this.parent(i)], this.heap[i]
	}
}
func (this *KthLargest) Pop() int {
	if this.Length() < 2 {
		return -1
	}

	res := this.heap[1]
	this.heap[1] = this.heap[this.Length()-1]
	this.heap = this.heap[:this.Length()-1]

	for i := 1; 2 * i < this.Length(); {
		l := this.left(i)	
		r := this.right(i)	
		largest := i

		if l < this.Length() && this.heap[l] > this.heap[largest] {
			largest = l
		}
		if r < this.Length() && this.heap[r] > this.heap[largest] {
			largest = r
		}
		if largest == i {
			break
		}

		this.heap[i], this.heap[largest] = this.heap[largest], this.heap[i]
		i = largest
	}

	return res
}
func (this *KthLargest) Length() int {
	return len(this.heap)
}
func (this *KthLargest) parent(i int) int {
	return i / 2
}
func (this *KthLargest) left(i int) int {
	return 2 * i
}
func (this *KthLargest) right(i int) int {
	return 2 * i + 1	
}