package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := piscine.Any(piscine.IsAlpha, a1)
	result2 := piscine.Any(piscine.IsAlpha, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
