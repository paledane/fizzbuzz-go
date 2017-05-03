package fizzbuzz

import "strconv"

func DefaultHandler(next func(int) string) func(int) string {
	return func(i int) string {
		return strconv.Itoa(i)
	}
}
