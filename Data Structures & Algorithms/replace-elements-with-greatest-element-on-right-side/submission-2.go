func replaceElements(arr []int) []int {
	n := len(arr)
	max_arr := make([]int, n)
	max_num := -(math.MaxInt)

	for i := (n-1); i >= 0; i-- {
		if arr[i] > max_num {
			max_num = arr[i]
		}
		max_arr[i] = max_num
	}

	for i := 0; i < n; i++ {
		if i == (n-1) {
			arr[i] = -1
			break
		}
		arr[i] = max_arr[i+1]
	}

	return arr
}
