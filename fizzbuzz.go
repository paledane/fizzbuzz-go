package fizzbuzz

//IsFizz returns true if value is divisable by 3
func IsFizz(value int) bool {
	return value%3 == 0
}

//IsBuzz returns true if value is divisable by 3
func IsBuzz(value int) bool {
	return value%5 == 0
}
