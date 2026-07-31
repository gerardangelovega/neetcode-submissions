// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func QuickSort(pairs []Pair) []Pair {
	qSort(pairs, 0, len(pairs) - 1)
	return pairs
}

func qSort(pairs []Pair, l, r int) {
	if l < r {
		i, k := l, l
		for i < r {
			if pairs[i].Key < pairs[r].Key {
				pairs[i], pairs[k] = pairs[k], pairs[i]
				k++
			}
			i++
		}

		pairs[r], pairs[k] = pairs[k], pairs[r]

		qSort(pairs, l, k - 1)
		qSort(pairs, k + 1, r)
	}
}
