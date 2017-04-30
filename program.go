package main

import (
	"fmt"

	"github.com/paledane/fizzbuzz-go/fizzbuzz"
)

func main() {

	for i := 1; i <= 100; i++ {
		if fizzbuzz.IsFizzBuzz(i) {
			fmt.Printf("FizzBuzz\n")
		} else if fizzbuzz.IsFizz(i) {
			fmt.Printf("Fizz\n")
		} else if fizzbuzz.IsBuzz(i) {
			fmt.Printf("Buzz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}

}
