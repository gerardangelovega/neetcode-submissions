type MinStack struct {
	arr []int
	minimum []int
}

func Constructor() MinStack {
	return MinStack{
		arr: make([]int, 0, 4),
		minimum: make([]int, 0, 4),
	}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)

	if len(this.minimum) == 0 {
		this.minimum = append(this.minimum, val)
	} else {
		n, m := len(this.arr), len(this.minimum)
		this.minimum = append(
			this.minimum, 
			min(this.arr[n-1], this.minimum[m-1]),
		)
	}
}

func (this *MinStack) Pop() {
	n, m := len(this.arr), len(this.minimum)
	this.arr = this.arr[:n-1]
	this.minimum = this.minimum[:m-1]
}

func (this *MinStack) Top() int {
	n := len(this.arr)
	val := this.arr[n-1]
	return val
}

func (this *MinStack) GetMin() int {
	m := len(this.minimum)
	val := this.minimum[m-1]
	return val
}
