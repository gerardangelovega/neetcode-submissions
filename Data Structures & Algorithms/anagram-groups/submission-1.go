func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]uint8][]string)
	res := make([][]string, 0, 0)

	for _, str := range strs {
		key := [26]uint8{}

		for i := 0; i < len(str); i++ {
			key[str[i]-97]++
		}

		anagrams[key] = append(anagrams[key], str)
	}

	for _, val := range anagrams {
		res = append(res, val)
	}

	return res
}
