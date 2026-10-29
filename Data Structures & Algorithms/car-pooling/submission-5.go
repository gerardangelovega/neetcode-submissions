func carPooling(trips [][]int, capacity int) bool {
	paths := make([]int, 1001)
	for _, trip := range trips {
		paths[trip[1]] = paths[trip[1]] + trip[0]
		paths[trip[2]] = paths[trip[2]] - trip[0]
	}

	sum := 0
	for i, _ := range paths {
		sum = sum + paths[i]
		paths[i] = sum

		if paths[i] > capacity { return false }
	}

	return true
}
