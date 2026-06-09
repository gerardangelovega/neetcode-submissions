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
	var meta_length strings.Builder
	length := 0
	reading_metadata := false
	reading_string := false

	for i, r := range encoded {
		if unicode.IsNumber(r) {
			meta_length.WriteRune(r)
			reading_metadata = true
		}
		if r == ' ' {
			meta_length.Reset()
			reading_metadata = false
		}
		if r == '#' && reading_metadata {
			length, _ = strconv.Atoi(meta_length.String())
			meta_length.Reset()
			reading_metadata = false
			reading_string = true
		}
		if reading_string {
			strs = append(strs, encoded[(i + 1):(i + length + 1)])
			reading_string = false
		}
	}

	return strs
}