package main

import (
	"fmt"

	"github.com/paledane/fizzbuzz-go/fizzbuzz"
)

func main() {

	for i := 1; i <= 100; i++ {
		fmt.Printf("%q\n", fizzbuzz.Calc(i))
	}

}
