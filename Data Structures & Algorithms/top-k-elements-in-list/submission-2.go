func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	max_frequency := 0

	for _, num := range nums {
		frequency[num]++
		if frequency[num] > max_frequency {
			max_frequency = frequency[num]
		}
	}

	bucket := make([][]int, max_frequency + 1)

	for key, val := range frequency {
		bucket[val] = append(bucket[val], key)
	}

	top_frequent := make([]int, 0, k)

	outer:
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}

		for j := len(bucket[i]) - 1; j >= 0; j-- {
			if len(top_frequent) >= k {
				break outer
			}
			top_frequent = append(top_frequent, bucket[i][j])
		}
	}

	return top_frequent
}
