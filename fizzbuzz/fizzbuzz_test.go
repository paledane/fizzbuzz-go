package fizzbuzz

import "testing"

func Test1Is1(t *testing.T) {
	want := "1"
	in := 1

	result := Calc(in)

	if result != want {
		t.Errorf("Calc(%d) == %v, want %v", in, result, want)
	}
}
