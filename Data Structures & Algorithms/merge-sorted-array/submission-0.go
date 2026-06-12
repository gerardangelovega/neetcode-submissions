func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m + n - 1, n - 1; j >= 0; i, j = i - 1, j - 1 {
		nums1[i] = nums2[j]
	}

	fmt.Println(nums1)

	mergeSort(nums1, 0, m + n - 1)
}

func mergeSort(arr []int, start, end int) {
	if (end - start + 1) <= 1 {
		return
	}

	mid := (end + start) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	mergeFunc(arr, start, mid, end)
}

func mergeFunc(arr []int, start, mid, end int) {
	n_left  := mid - start + 1
	n_right := end - mid

	left := make([]int, n_left)
	copy(left, arr[start:mid+1])

	right := make([]int, n_right)
	copy(right, arr[mid+1:end+1])

	l, r, i := 0, 0, start
	for l < n_left && r < n_right {
		if left[l] <= right[r] {
			arr[i] = left[l]
			l++
		} else {
			arr[i] = right[r]
			r++
		}
		i++
	}

	for l < n_left {
		arr[i] = left[l]
		l++
		i++
	}
	for r < n_right {
		arr[i] = right[r]
		r++
		i++
	}
}
