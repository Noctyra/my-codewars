package reversed_string

func Reversed(word string) string {
	result := ""
	for _, c := range word {
		result = string(c) + result
	}
	return result
}
