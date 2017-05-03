package fizzbuzz

import "strconv"

func DefaultHandler(next func(int) string) func(int) string {
	return func(i int) string {
		return strconv.Itoa(i)
	}
}

func FizzHandler(next func(int) string) func(int) string {
	return func(i int) string {
		if i%3 == 0 {
			return "Fizz"
		}
		return strconv.Itoa(i)
	}
}
