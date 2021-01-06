package fizz_buzz

import (
	"strconv"
	"strings"
)

func FizzBuzz(int1, int2, limit int, str1, str2 string) string {
	var outputSlice []string
	var r string

	for i := 1; i <= limit; i++ {
		r = ""
		if i%int1 == 0 {
			r += str1
		}

		if i%int2 == 0 {
			r += str2
		}

		if r == "" {
			r = strconv.FormatInt(int64(i), 10)
		}

		outputSlice = append(outputSlice, r)
	}

	return strings.Join(outputSlice[:], ",")
}
