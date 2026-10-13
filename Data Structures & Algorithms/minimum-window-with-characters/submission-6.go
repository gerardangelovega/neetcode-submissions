func minWindow(s string, t string) string {
	if len(t) == 0 { return "" }

	// tFrequency; keeps track of how many of each character is in t
	// wFrequency; keeps track of how many of each character is in the current window
	tFrequency, wFrequency := make(map[byte]int), make(map[byte]int)
	for _, c := range t { tFrequency[byte(c)]++ }

	// have; keeps track how many characters of t is in the current window
	// need; how many characters of t that the window needs to form a valid substring
	have, need := 0, len(tFrequency)

	// res; stores the range of the substring
	// resLen; stores the length of the substring derived from the range and is used for comparisons
	res, resLen := []int{ -1, -1 }, math.MaxInt

	// l; left pointer of the window (shrink)
	// r; right pointer of the window (grow)
	l, r := 0, 0
	for ; r < len(s); r++ {
		// increment the frequency of character in the wFrequency whenever the right pointer 
		// encounters a character in the string
		wFrequency[s[r]]++

		// if the character exists in t and the frequency of the character in both t string and the 
		// window match, then increment have since we have enough of that character
		if tFrequency[s[r]] > 0 && tFrequency[s[r]] == wFrequency[s[r]] {
			have++
		}

		// if have is equal to need, meaning all characters of t is present in window, then:
		// - we check if the current window forms the minimum valid substring
		// - shrink window from the left
		// - decrement have if the frequency of a character in the window is less than the frequency
		//	 of the character in t
		// - repeat until have != need
		for ; have == need; l++ {
			// check if smaller, if so, update
			if (r - l + 1) < resLen {
				res = []int{ l, r }
				resLen = r - l + 1
			}

			// shrink window from the left
			wFrequency[s[l]]--
			// if character appears less in window compared to string t,
			// decrement have since we don't meet the minimum occurences of the char in t
			if tFrequency[s[l]] > 0 && wFrequency[s[l]] < tFrequency[s[l]] {
				have--
			}
		}
	}

	// if res is still { -1, -1 }, it means no valid susbtring was found so return ""
	if res[0] == -1 {
		return ""
	}

	// slice the string with res
	return s[res[0]:res[1]+1]
}
