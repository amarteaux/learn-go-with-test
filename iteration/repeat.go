package iteration

func Repeat(char string, rep int) string {
	var result string
	for i := 0; i < rep; i++ {
		result += char
	}
	return result
}
