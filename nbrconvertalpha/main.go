package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := 0
	for _, arg := range args {
		number := 0
		for _, v := range arg {
			number = number*10 + int(v-'0')
		}
		if arg == "--upper" {
			upper = 1
		}
		if arg == "--upper" || number > 26 {
			z01.PrintRune(' ')
			continue
		}
		z01.PrintRune(rune((number + 'a' - 1 - upper*32)))
	}
	z01.PrintRune('\n')
}
