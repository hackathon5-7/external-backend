package concatenation

func ConcatenationString(data ...string) string {
	var result string

	for _, v := range data {
		result += v
	}

	return result
}
