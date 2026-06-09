type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteString("#")
		builder.WriteString(str)
	}

	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	strs := make([]string, 0)

	for start, end := 0, 0; end < len(encoded); {
		if encoded[end] != '#' { 
			end++
			continue 
		}

		length, _ := strconv.Atoi(encoded[start:end])
		start = end + 1
		end = start + length

		strs = append(strs, encoded[start:end])

		start = end
	}

	return strs
}