package toLowerCase

func ToLowerCase(str string) string {
	cs := []rune(str)
	for i := 0; i < len(cs); i++ {
		if 65 <= cs[i] && cs[i] <= 90 {
			cs[i] = cs[i] + 32
		}
	}
	return string(cs)
}
