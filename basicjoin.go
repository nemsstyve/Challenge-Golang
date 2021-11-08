package piscine

func BasicJoin(elems []string) string {
	den := ""
	for _, s := range elems {
		den += s
	}
	return den
}
