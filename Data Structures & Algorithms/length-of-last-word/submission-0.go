func lengthOfLastWord(s string) int {
	str  := strings.Trim(s, " ")	
	strs := strings.Split(str, " ")

	return len(strs[len(strs)-1])
}
