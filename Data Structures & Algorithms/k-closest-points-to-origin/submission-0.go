func kClosest(points [][]int, k int) [][]int {
	quickSort(points, 0, len(points) - 1)
	return points[:k]
}

func quickSort(arr [][]int, start, end int) {
	if (end - start) + 1 <= 1 {
		return
	}

	left := start
	for i := start; i < end; i++ {
		if distance(arr[i]) < distance(arr[end]) {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[end], arr[left] = arr[left], arr[end]

	quickSort(arr, start,  left-1)
	quickSort(arr, left+1, end)
}

func distance(point []int) float64 {
	return math.Sqrt(float64(point[0] * point[0] + point[1] * point[1]))
}
