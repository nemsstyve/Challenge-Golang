package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	length := 0
	for index := range params {
		length = index
	}
	for i := length; i >= 1; i-- {
		for _, w := range params[i] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
