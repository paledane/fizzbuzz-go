package main

import (
	"fmt"

	"github.com/paledane/fizzbuzz-go/fizzbuzz"
)

func main() {

	chain := fizzbuzz.FizzBuzzHandler(
		fizzbuzz.FizzHandler(
			fizzbuzz.BuzzHandler(
				fizzbuzz.DefaultHandler(nil))))

	for i := 1; i <= 100; i++ {
		fmt.Printf("%q\n", chain(i))
	}

}
