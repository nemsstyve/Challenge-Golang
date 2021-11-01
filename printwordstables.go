package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, c := range a {
		for _, j := range c {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
