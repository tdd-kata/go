package loop

const repeatCount = 5

// Repeat returns character repeated 5 times.
func Repeat(character string) string {
	var repeated string

	// while, do, until 키워드는 없다.
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
