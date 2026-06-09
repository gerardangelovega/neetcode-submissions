type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var builder strings.Builder

	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteString("#")
		builder.WriteString(str)
		builder.WriteString(" ")
	}

	return strings.Trim(builder.String(), " ")
}

func (s *Solution) Decode(encoded string) []string {
	if len(encoded) == 0 {
		return []string{}
	}

	strs := make([]string, 0)

	for start, end := 0, 0; end < len(encoded); end++ {
		if encoded[end] != '#' { continue }

		length, _ := strconv.Atoi(encoded[start:end])
		start, end = end + 1, end + 1 + length
		strs = append(strs, encoded[start:end])

		start, end = end + 1, end + 1
	}

	return strs
}