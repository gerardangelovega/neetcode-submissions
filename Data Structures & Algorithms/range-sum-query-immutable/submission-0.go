type NumArray struct {
    arr []int
	pre []int
}


func Constructor(nums []int) NumArray {
    arr := make([]int, len(nums));
    pre := make([]int, len(nums));
	copy(arr, nums)
	copy(pre, nums)

	for i := 1; i < len(pre); i++ {
		pre[i] = pre[i-1] + pre[i]
	}

	return NumArray {
		arr: arr,
		pre: pre,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 { return this.pre[right] }
	return this.pre[right] - this.pre[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */