// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
	return mySort(pairs, 0, len(pairs) - 1)
}

func mySort(pairs []Pair, start, end int) []Pair {
	if (end - start + 1) <= 1 {
		return pairs
	}

	mid := (end + start) / 2
	mySort(pairs, start, mid)
	mySort(pairs, mid+1, end)

	merge(pairs, start, mid, end)

	return pairs
}

func merge(pairs []Pair, start, mid, end int) {
	n_left  := mid - start + 1
	n_right := end - mid

	left := make([]Pair, n_left)
	copy(left, pairs[start:mid+1])

	right := make([]Pair, n_right)
	copy(right, pairs[mid+1:end+1])

	l, r, i := 0, 0, start
	for l < n_left && r < n_right {
		if left[l].Key <= right[r].Key {
			pairs[i] = left[l]
			l++
		} else {
			pairs[i] = right[r]
			r++
		}
		i++
	}

	for l < n_left {
		pairs[i] = left[l]
		l++
		i++
	}
	for r < n_right {
		pairs[i] = right[r]
		r++
		i++
	}
}
