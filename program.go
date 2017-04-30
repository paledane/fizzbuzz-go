package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {
		if IsFizzBuzz(i) {
			fmt.Printf("FizzBuzz\n")
		} else if IsFizz(i) {
			fmt.Printf("Fizz\n")
		} else if IsBuzz(i) {
			fmt.Printf("Buzz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}

}
