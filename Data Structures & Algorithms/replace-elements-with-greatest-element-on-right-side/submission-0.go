func replaceElements(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		if i == (n - 1) {
			arr[i] = -1
			break
		}

		greatest := -(math.MaxInt)

		for j := i + 1; j < n; j++ {
			greatest = max(greatest, arr[j])
		}

		arr[i] = greatest
	}

	return arr
}
