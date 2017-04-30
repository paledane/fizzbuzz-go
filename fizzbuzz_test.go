package main

import "testing"

func TestIsFizz(t *testing.T) {
	in := 3
	want := true
	result := IsFizz(in)

	if result != want {
		t.Errorf("IsFizz(%d) == %t, want %t", in, result, want)
	}
}

func TestIsNotFizz(t *testing.T) {
	in := 2
	want := false
	result := IsFizz(in)

	if result != want {
		t.Errorf("IsFizz(%d) == %t, want %t", in, result, want)
	}
}

func TestIsBuzz(t *testing.T) {
	in := 5
	want := true
	result := IsBuzz(in)

	if result != want {
		t.Errorf("IsBuzz(%d) == %t, want %t", in, result, want)
	}
}

func TestIsNotBuzz(t *testing.T) {
	in := 3
	want := false
	result := IsBuzz(in)

	if result != want {
		t.Errorf("IsBuzz(%d) == %t, want %t", in, result, want)
	}
}

func TestIsFizzBuzz(t *testing.T) {
	in := 15
	want := true
	result := IsFizzBuzz(in)

	if result != want {
		t.Errorf("IsFizzBuzz(%d) == %t, want %t", in, result, want)
	}
}
