// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func QuickSort(pairs []Pair) []Pair {
	quickSort(pairs, 0, len(pairs)-1)
	return pairs
}

func quickSort(pairs []Pair, start, end int) {
	if (end - start) + 1 <= 1 {
		return
	}

	left := start
	for i := start; i < end; i++ {
		if pairs[i].Key < pairs[end].Key {
			pairs[i], pairs[left] = pairs[left], pairs[i]
			left++
		}
	}

	pairs[left], pairs[end] = pairs[end], pairs[left]

	quickSort(pairs, start,  left-1)
	quickSort(pairs, left+1, end)
}
