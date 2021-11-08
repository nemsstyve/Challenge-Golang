package piscine

func FindNextPrime(number int) int {
	for !IsPrime(number) {
		number++
	}
	return number

}
