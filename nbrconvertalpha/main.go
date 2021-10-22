package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	len := 0
	flag := false
	for i := range arg {
		len = i
	}
	for _, val := range arg {
		if val == "--upper" {
			flag = true
		}
	}
	if flag {
		for i := 0; i <= len; i++ {
			z01.PrintRune(ToUpper(ConvertToLetter(arg[i])))
		}
		z01.PrintRune(10)
	} else if len == 0 {
		z01.PrintRune(10)
	} else {
		for i := 0; i <= len; i++ {
			z01.PrintRune(ConvertToLetter(arg[i]))
		}
		z01.PrintRune(10)
	}
}
func ConvertToLetter(s string) rune {
	numb := 0
	for _, i := range s {
		count := 0
		for k := '0'; k < i; k++ {
			count++
		}
		numb = numb*10 + count
	}
	numb = numb + 96
	if numb >= 123 {
		return ' '
	}
	return rune(numb)
}
func ToUpper(r rune) rune {
	if r-32 < 65 {
		return r
	}
	return r - 32
}
