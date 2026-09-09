func checkInclusion(s1 string, s2 string) bool {
	n := len(s2)
	m := make(map[rune]int)

	for _, r := range s1 {
		m[r]++
	}

	l, r := 0, 0

	for r < n {
		c := rune(s2[r])
		if _, e := m[c]; !e {
			if check(m) == 0 { return true }
			for l <= r {
				c2 := rune(s2[l])
				if _, e := m[rune(c2)]; e { m[c2]++ }
				l++
			}
			r++
			continue
		}

		if _, e := m[c]; e {
			if check(m) == 0 { return true }
			m[c]--
		} 
		
		if v, e := m[c]; e && v < 0 {
			if check(m) == 0 { return true }
			for l <= r {
				c2 := rune(s2[l])
				if _, e := m[c2]; e {
					m[c2]++
					if c2 == c {
						l++
						break
					}
				}
				l++
			}
		}
		fmt.Println("\t",m[c], check(m), c)
		r++
	}
	return check(m) == 0
}

func check(m map[rune]int) int {
	res := 0
	for _, v := range m {
		if v < 0 {
			return -1
		}
		res = res + v
	}
	return res
}

// Create hash map composed of the characters in s1 and a boolean indicating if they have been encountered
// Use two pointers
// Move the right pointer until all characters of the permutation occurs once
// Move the left pointer until it meets the first character of the permutation
// If the right pointer encounters any character not a part of s1, move left to right
// If the right pointer encounters a char of s1 twice, move left to right