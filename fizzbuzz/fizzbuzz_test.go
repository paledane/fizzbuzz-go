package fizzbuzz

import "testing"

func TestDefaultHandler(t *testing.T) {
	sut := DefaultHandler(nil)
	in := 1
	want := "1"

	if result := sut(in); result != want {
		t.Errorf("Calc(%d) == %q, want %q", in, result, want)
	}
}
