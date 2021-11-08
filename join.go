package piscine

func Join(strs []string, sep string) string {
	dev := ""
	first := true
	for _, s := range strs {
		if first {
			dev = s
			first = false
		} else {
			dev += sep + s
		}
	}
	return dev
}
