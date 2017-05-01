package fizzbuzz

import (
	"fmt"
)

type Fizzbuzzer interface {
	Calc() string
}

type Fizzbuzz int32

// Calc returns the specified number. But for multiples of three
// “Fizz” is returned instead of the number and for the multiples
// of five print “Buzz” is returned. For numbers which are multiples
// of both three and five “FizzBuzz” is returned.
func (i Fizzbuzz) Calc() string {
	if i%15 == 0 {
		return "FizzBuzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	if i%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%v", i)

}
