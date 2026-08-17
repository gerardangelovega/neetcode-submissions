func longestCommonSubsequence(text1 string, text2 string) int {
	return memoize(text1, text2, 0, 0, make(map[[2]int]int))
}

func memoize(text1 string, text2 string, i, j int, cache map[[2]int]int) int {
	if i == len(text1) || j == len(text2) {
		return 0
	}
	key := [2]int{i, j}
	if text1[i] == text2[j] {
		return 1 + memoize(text1, text2, i + 1, j + 1, cache)
	}
	if val, exists := cache[key]; exists {
		return val
	}
	cache[key] = max(memoize(text1, text2, i + 1, j, cache), memoize(text1, text2, i, j + 1, cache))
	return cache[key]
}