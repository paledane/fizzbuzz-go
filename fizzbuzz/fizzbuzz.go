package fizzbuzz

import (
	"strconv"
)

type Handler interface {
	Calc(i int) string
}

type FizzHandler struct {
	Next Handler
}

type BuzzHandler struct {
	Next Handler
}

type FizzBuzzHandler struct {
	Next Handler
}

type DefaultHandler struct {
}

func (handler *FizzHandler) Calc(i int) string {
	if i%3 == 0 {
		return "Fizz"
	}
	return handler.Next.Calc(i)
}

func (handler *BuzzHandler) Calc(i int) string {
	if i%5 == 0 {
		return "Buzz"
	}
	return handler.Next.Calc(i)
}

func (handler *FizzBuzzHandler) Calc(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	}
	return handler.Next.Calc(i)
}

func (handler *DefaultHandler) Calc(i int) string {
	return strconv.Itoa(i)
}
