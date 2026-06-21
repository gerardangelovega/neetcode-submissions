func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums) - 1)
	return nums
}

func quickSort(arr []int, start, end int) {
	if (end - start + 1) <= 1 {
		return
	}

	left := start
	for i := start; i < end; i++ {
		if arr[i] < arr[end] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[end] = arr[end], arr[left]

	quickSort(arr, start, left-1)
	quickSort(arr, left+1, end)
}