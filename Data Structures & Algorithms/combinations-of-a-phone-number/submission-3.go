func letterCombinations(digits string) []string {
	n := len(digits)
	dtl := map[byte][]byte {
		'1': []byte{},
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	combinations := make([]string, 0, int(math.Pow(4, float64(n))))
	combination := make([]byte, 0, n)

	var generate func(int)
	generate = func(i int) {
		if i > n-1 {
			if len(combination) > 0 {
				combinations = append(combinations, string(combination))
			}
			return
		}
		digit := digits[i]
		letters := dtl[digit]
		for _, letter := range letters {
			combination = append(combination, letter)
			generate(i + 1)
			combination = combination[:len(combination)-1]
		}
	}
	generate(0);

	return combinations
}

func pow(base, exp int) int {
	total := 1
	for range exp {
		total = total * base
	}
	return total
}