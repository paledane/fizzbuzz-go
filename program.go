package main

import (
	"fmt"

	"github.com/paledane/fizzbuzz-go/fizzbuzz"
)

func main() {

	chain := &fizzbuzz.FizzBuzzHandler{
		Next: &fizzbuzz.FizzHandler{
			Next: &fizzbuzz.BuzzHandler{
				Next: &fizzbuzz.DefaultHandler{},
			},
		},
	}

	for i := 1; i <= 100; i++ {
		fmt.Printf("%q\n", chain.Calc(i))
	}

}
