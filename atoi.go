package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	for i, r := range s {
		if i == 0 {
			if r == '-' {
				sign = -1
				continue
			}
			if r == '+' {
				continue
			}
		}
		if r < '0' || r > '9' {
			return 0
		}
		for i := 0; i <= 9; i += 1 {
			if rune(i) == r-'0' {
				result *= 10
				result += i
				break
			}
		}
	}
	return result * sign
}
