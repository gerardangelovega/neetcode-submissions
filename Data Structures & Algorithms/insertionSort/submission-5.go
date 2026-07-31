// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	n := len(pairs)
	res := [][]Pair{}

	for i := 0; i < n; i++ {
		j := i - 1
		for j >= 0 && pairs[j + 1].Key < pairs[j].Key {
			pairs[j], pairs[j + 1] = pairs[j + 1], pairs[j]
			j--
		}

		clone := make([]Pair, n)
		copy(clone, pairs)
		res = append(res, clone)
	}

	return res
}
