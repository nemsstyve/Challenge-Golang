package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0
	for range os.Args {
		count++
	}
	for i := 1; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if os.Args[i] > os.Args[j] {
				temp := os.Args[j]
				os.Args[j] = os.Args[i]
				os.Args[i] = temp
			}
		}
	}
	i := 1
	for i < count {
		out := []rune(os.Args[i])
		for _, j := range out {
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
		i++
	}
}
