package piscine

func Map(f func(int) bool, a []int) []bool {
	r := []bool{}
	for _, s := range a {
		r = append(r, f(s))
	}
	return r
}
