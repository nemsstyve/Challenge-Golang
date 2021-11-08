package piscine

func StrRev(s string) string {
	sty := []rune(s)
	for i, j := 0, len(sty)-1; i < j; i, j = i+1, j-1 {
		sty[i], sty[j] = sty[j], sty[i]
	}
	return string(sty)
}
