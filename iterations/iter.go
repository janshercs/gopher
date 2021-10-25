package iteration // this is the name of this package

// Takes a string and an integer and returns the same string repeated by the integer given.
func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
