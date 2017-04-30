package fizzbuzz

import "testing"

func Test1Is1(t *testing.T) {
	want := "1"
	in := 1

	if result := Calc(in); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func Test13Is13(t *testing.T) {
	want := "13"
	in := 13

	if result := Calc(in); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func Test3IsFizz(t *testing.T) {
	want := "Fizz"
	in := 3

	if result := Calc(in); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func Test5IsBuzz(t *testing.T) {
	want := "Buzz"
	in := 5

	if result := Calc(in); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}
