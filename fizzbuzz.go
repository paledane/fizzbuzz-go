package fizzbuzz

// IsFizz returns true if value is divisable by 3
func IsFizz(value int) bool {
	return value%3 == 0
}

// IsBuzz returns true if value is divisable by 5
func IsBuzz(value int) bool {
	return value%5 == 0
}

// IsFizzBuzz returns true if value is divisable by 3 and 5
func IsFizzBuzz(value int) bool {
	return value%15 == 0
}
