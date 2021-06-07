package iterations

func Repeat(character string, repetitionsQtty int) string {
	var repetitions string
	for i := 0; i < repetitionsQtty; i++ {
		repetitions += character
	}
	return repetitions
}
