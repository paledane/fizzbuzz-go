package main

import (
	"fmt"

	"github.com/paledane/fizzbuzz-go/fizzbuzz"
)

func main() {
	var fb fizzbuzz.Fizzbuzzer
	for i := 1; i <= 100; i++ {
		fb = fizzbuzz.Fizzbuzz(i)
		fmt.Printf("%q\n", fb.Calc())
	}

}
