package piscine

func BasicAtoi(s string) int {
	x := 0
	for _, c := range s {
		che := 0
		for i := '1'; i <= c; i++ {
			che++
		}
		x = x*10 + che
	}
	return x
}
