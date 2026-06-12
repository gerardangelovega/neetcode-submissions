// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	n := len(pairs)
	res := make([][]Pair, 0, n)
	clone := make([]Pair, n)

	if n > 0 {
		copy(clone, pairs)
		res = append(res, clone)
	}

	for i := 1; i < n; i++ {
		if pairs[i].Key < pairs[i-1].Key {
			for j := i; j > 0; j-- {
				if pairs[j].Key >= pairs[j-1].Key {
					break
				}
				pairs[j], pairs[j-1] = pairs[j-1], pairs[j]
			}
		}
		clone = make([]Pair, n)
		copy(clone, pairs)
		res = append(res, clone)
	}

	return res
}
