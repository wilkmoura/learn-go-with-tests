package iteration

func Repeat(character string, numRepeatition int) string {
	var repeated string
	for i := 0; i < numRepeatition; i++ {
		repeated = repeated + character
	}
	return repeated
}
