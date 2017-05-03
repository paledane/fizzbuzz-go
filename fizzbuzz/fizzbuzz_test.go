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

func fakeHandler(next func(int) string) func(int) string {
	return func(i int) string {
		return "This is the end"
	}
}

func TestCanReachEndOfChain(t *testing.T) {
	sut := FizzHandler(fakeHandler(nil))
	in := 1
	want := "This is the end"

	if result := sut(in); result != want {
		t.Errorf("func(%d) == %q, want %q", in, result, want)
	}
}

func TestCanStopAtFirstHandlerInChain(t *testing.T) {
	sut := FizzHandler(fakeHandler(nil))
	in := 3
	want := "Fizz"

	if result := sut(in); result != want {
		t.Errorf("func(%d) == %q, want %q", in, result, want)
	}
}

func TestBuzzHandler(t *testing.T) {
	sut := BuzzHandler(nil)
	in := 5
	want := "Buzz"

	if result := sut(in); result != want {
		t.Errorf("func(%d) == %q, want %q", in, result, want)
	}
}

func TestFizzBuzzHandler(t *testing.T) {
	sut := FizzBuzzHandler(nil)
	in := 15
	want := "FizzBuzz"

	if result := sut(in); result != want {
		t.Errorf("func(%d) == %q, want %q", in, result, want)
	}
}
