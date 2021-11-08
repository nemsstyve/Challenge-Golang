package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(number int) {
	var decimal_slice [10]int
	if number == 0 {
		z01.PrintRune('0')
		return
	}
	for number != 0 {
		decimal_slice[number%10]++
		number /= 10
	}
	for i := 0; i < 10; i++ {
		for decimal_slice[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			decimal_slice[i]--
		}
	}
}
