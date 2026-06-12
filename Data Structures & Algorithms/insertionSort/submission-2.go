// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	n := len(pairs)
	res := make([][]Pair, 0, n)

	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if pairs[j].Key >= pairs[j-1].Key {
				break
			}
			pairs[j], pairs[j-1] = pairs[j-1], pairs[j]
		}
		clone := make([]Pair, n)
		copy(clone, pairs)
		res = append(res, clone)
	}

	return res
}