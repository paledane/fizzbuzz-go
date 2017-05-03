package fizzbuzz

import "testing"

func TestDefaultHandler(t *testing.T) {
	sut := DefaultHandler(nil)
	in := 1
	want := "1"

	if result := sut(in); result != want {
		t.Errorf("func(%d) == %q, want %q", in, result, want)
	}
}

func TestFizzHandler(t *testing.T) {
	sut := FizzHandler(nil)
	in := 3
	want := "Fizz"

	if result := sut(in); result != want {
		t.Errorf("func(%d) == %q, want %q", in, result, want)
	}
}
