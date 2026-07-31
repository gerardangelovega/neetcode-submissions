// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	res := make([][]Pair, 0, len(pairs))

	if len(pairs) == 0 {
		return res
	}

	for i := 1; i < len(pairs); i++ {
		clone := make([]Pair, len(pairs))
		copy(clone, pairs)
		res = append(res, clone)
		
		j := i - 1
		for j >= 0 && pairs[j + 1].Key < pairs[j].Key {
			pairs[j], pairs[j + 1] = pairs[j + 1], pairs[j]
			j--
		}
	}

	clone := make([]Pair, len(pairs))
	copy(clone, pairs)
	res = append(res, clone)

	return res
}
