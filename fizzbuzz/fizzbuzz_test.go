package fizzbuzz

import "testing"

func Test1Is1(t *testing.T) {
	want := "1"
	in := 1

	result := Calc(in)

	if result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func Test13Is13(t *testing.T) {
	want := "13"
	in := 13

	result := Calc(in)

	if result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func Test3IsFizz(t *testing.T) {
	want := "Fizz"
	in := 3

	result := Calc(in)

	if result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}
