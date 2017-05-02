package fizzbuzz

import "testing"

func TestDefaultHandler(t *testing.T) {
	in := 1
	handler := DefaultHandler{}
	want := "1"

	if result := handler.Calc(in); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func TestFizzHandler(t *testing.T) {
	in := 3
	handler := FizzHandler{}
	want := "Fizz"

	if result := handler.Calc(in); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func TestBuzzHandler(t *testing.T) {
	in := 5
	handler := BuzzHandler{}
	want := "Buzz"

	if result := handler.Calc(in); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func TestFizzBuzzHandler(t *testing.T) {
	in := 15
	handler := FizzBuzzHandler{}
	want := "FizzBuzz"

	if result := handler.Calc(in); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}
