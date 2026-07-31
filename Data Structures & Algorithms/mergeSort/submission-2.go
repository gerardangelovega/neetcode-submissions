// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
	return mSort(pairs, 0, len(pairs) - 1)
}

func mSort(pairs []Pair, l, r int) []Pair {
	if l < r {
		m := (l + r) / 2

		mSort(pairs, l, m)
		mSort(pairs, m + 1, r)
		merge(pairs, l, m, r)
	}

	return pairs
}

func merge(pairs []Pair, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	left_arr := make([]Pair, n1)
	right_arr := make([]Pair, n2)

	copy(left_arr, pairs[l:m+1])
	copy(right_arr, pairs[m+1:r+1])

	// i is left_arr pointer
	// j is right_arr pointer
	// k is pairs pointer
	i, j, k := 0, 0, l

	for i < n1 && j < n2 {
		if left_arr[i].Key <= right_arr[j].Key {
			pairs[k] = left_arr[i]
			i++
		} else {
			pairs[k] = right_arr[j]
			j++
		}
		k++
	}

	for i < n1 {
		pairs[k] = left_arr[i]
		i++
		k++
	}

	for j < n2 {
		pairs[k] = right_arr[j]
		j++
		k++
	}
}