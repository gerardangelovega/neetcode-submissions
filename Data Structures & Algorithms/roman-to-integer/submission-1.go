func romanToInt(s string) int {
	lookup := map[byte]int{
		'I': 1, 
		'V': 5, 
		'X': 10, 
		'L': 50, 
		'C': 100, 
		'D': 500, 
		'M': 1000,
	}

	res := 0
	tmp := 0
	for i := 0; i < len(s); i++ {
		fmt.Println(res, tmp)
		curr := lookup[s[i]]
		fmt.Println(curr, tmp)
		if curr > tmp && tmp != 0 {
			res = res + (curr - tmp)
			tmp = 0
		} else if curr < tmp {
			res = res + tmp
			tmp = curr
		} else {
			tmp = tmp + curr
		}
	}
	res = res + tmp
	return res
}
