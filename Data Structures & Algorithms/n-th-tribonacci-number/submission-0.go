func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	cache := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		sum := 0
		for _, c := range cache { sum = sum + c }
		cache = cache[1:]
		cache = append(cache, sum)
	}
	return cache[2]
}
