package fizzbuzz

import "testing"

func TestIsFizz(t *testing.T) {
	in := 3
	want := true
	result := IsFizz(in)

	if result != want {
		t.Errorf("IsFizz(%v) == %t, want %t", in, result, want)
	}
}

func TestIsNotFizz(t *testing.T) {
	in := 2
	want := false
	result := IsFizz(in)

	if result != want {
		t.Errorf("IsFizz(%v) == %t, want %t", in, result, want)
	}
}

func TestIsBuzz(t *testing.T) {
	in := 5
	want := true
	result := IsBuzz(in)

	if result != want {
		t.Errorf("IsBuzz(%v) == %t, want %t", in, result, want)
	}
}

func TestIsNotBuzz(t *testing.T) {
	in := 3
	want := false
	result := IsBuzz(in)

	if result != want {
		t.Errorf("IsBuzz(%v) == %t, want %t", in, result, want)
	}
}
