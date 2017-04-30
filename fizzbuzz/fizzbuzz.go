package fizzbuzz

import (
	"strconv"
)

// Calc returns the specified number. But for multiples of three
// “Fizz” is returned instead of the number and for the multiples
// of five print “Buzz” is returned. For numbers which are multiples
// of both three and five “FizzBuzz” is returned.
func Calc(i int) string {
	if i%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(i)

}
