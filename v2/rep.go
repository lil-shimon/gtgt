package rep

func Repeat(word string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {
		repeated += word
	}
	return repeated
}
