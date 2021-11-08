package piscine

func check(s string) bool {
	for _, c := range s {
		if c == ' ' || c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func BasicAtoi2(str string) int {
	x := 0
	if check(str) == true {
		for _, c := range str {
			cnt := 0
			for i := '1'; i <= c; i++ {
				cnt++
			}
			x = x*10 + cnt
		}
	}
	return x
}
