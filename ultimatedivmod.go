package piscine

func UltimateDivMod(a *int, b *int) {
	tempA := *a / *b
	tempB := *a % *b
	*a = tempA
	*b = tempB
}
