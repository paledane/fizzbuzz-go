package fizzbuzz

import "strconv"

func DefaultHandler(next func(int) string) func(int) string {
	return func(i int) string {
		return strconv.Itoa(i)
	}
}

func FizzHandler(next func(int) string) func(int) string {
	handler := next
	return func(i int) string {
		if i%3 == 0 {
			return "Fizz"
		}
		return handler(i)
	}
}

func BuzzHandler(next func(int) string) func(int) string {
	handler := next
	return func(i int) string {
		if i%5 == 0 {
			return "Buzz"
		}
		return handler(i)
	}
}

func FizzBuzzHandler(next func(int) string) func(int) string {
	handler := next
	return func(i int) string {
		if i%15 == 0 {
			return "FizzBuzz"
		}
		return handler(i)
	}
}
