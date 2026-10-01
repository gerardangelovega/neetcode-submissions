func findClosestElements(arr []int, k int, x int) []int {
	res := make([]int, k)
	n := len(arr)
	l, r := 0, k-1
	diff := math.MaxInt

	for r < n {
		d := 0
		for _, num := range arr[l:r+1] { d = d + abs(x - num) }
		if d < diff {
			diff = d
			res = arr[l:r+1]
		}
		l++
		r++
	}
	return res
}

func abs(x int) int {
	if x < 0 { x = x * -1 }
	return x
}
