import "slices"

func merge(nums1 []int, m int, nums2 []int, n int) {
	left := slices.Clone(nums1[:m])
	right := nums2

	l, r, i := 0, 0, 0
	for l < m && r < n {
		if left[l] <= right[r] {
			nums1[i] = left[l]
			l++
		} else {
			nums1[i] = right[r]
			r++
		}
		i++
	}

	for l < m {
		nums1[i] = left[l]
		l++
		i++
	}
	for r < n {
		nums1[i] = right[r]
		r++
		i++
	}
}
