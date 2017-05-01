package fizzbuzz

import "testing"

func Test1Is1(t *testing.T) {
	var in Fizzbuzzer = Fizzbuzz(1)
	want := "1"

	if result := in.Calc(); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func Test13Is13(t *testing.T) {
	var in Fizzbuzzer = Fizzbuzz(13)
	want := "13"

	if result := in.Calc(); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func Test3IsFizz(t *testing.T) {
	var in Fizzbuzzer = Fizzbuzz(3)
	want := "Fizz"

	if result := in.Calc(); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func Test5IsBuzz(t *testing.T) {
	var in Fizzbuzzer = Fizzbuzz(5)
	want := "Buzz"

	if result := in.Calc(); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}

func Test15IsFizzBuzz(t *testing.T) {
	var in Fizzbuzzer = Fizzbuzz(15)
	want := "FizzBuzz"

	if result := in.Calc(); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}
