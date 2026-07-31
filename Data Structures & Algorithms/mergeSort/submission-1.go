// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
	return mSort(pairs, 0, len(pairs) - 1)
}

func mSort(pairs []Pair, left, right int) []Pair {
	if left < right {
		mid := (left + right) / 2

		mSort(pairs, left, mid)
		mSort(pairs, mid + 1, right)
		merge(pairs, left, mid, right)
	}

	return pairs
}

func merge(pairs []Pair, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	l_arr := make([]Pair, n1)
	r_arr := make([]Pair, n2)

	for i := 0; i < n1; i++ {
		l_arr[i] = pairs[left + i]
	}

	for i := 0; i < n2; i++ {
		r_arr[i] = pairs[mid + 1 + i]
	}

	i, j, k := 0, 0, left

	for i < n1 && j < n2 {
		if l_arr[i].Key <= r_arr[j].Key {
			pairs[k] = l_arr[i]
			i++
		} else {
			pairs[k] = r_arr[j]
			j++
		}
		k++
	}

	for i < n1 {
		pairs[k] = l_arr[i]
		i++
		k++
	}

	for j < n2 {
		pairs[k] = r_arr[j]
		j++
		k++
	}
}