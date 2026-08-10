import "slices"

func minEatingSpeed(piles []int, h int) int {
	left_k := 1
	right_k := slices.Max(piles)
	mid_k := (left_k + right_k) / 2
	
	min_k := math.MaxInt

	for left_k <= right_k {
		if check(piles, h, mid_k) {
			min_k = min(min_k, mid_k)
			right_k = mid_k - 1
		} else {
			left_k = mid_k + 1
		}
		mid_k = (left_k + right_k) / 2
	}

	return min_k
}

func check(piles []int, h, k int) bool {
	i := 0

	for _, pile := range piles {
		i = i + (pile + k - 1) / k
	}

	return i <= h
}
