package iteration

import "strings"

func Repeat(character string, times int) string {
	//var repeat string
	//for i := 0; i < times; i++ {
	//	repeat = repeat + character
	//}
	//return repeat

	if times < 0 {
		times = 0
	}

	return strings.Repeat(character, times)
}
