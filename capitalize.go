package piscine

func Capitalize(s string) string {
	pes := []rune(s)
	strlen := 0
	for i := range pes {
		strlen = i + 1
	}
	for i := 0; i < strlen; i++ {
		if i != 0 && (!((pes[i-1] >= 'a' && pes[i-1] <= 'z') || (pes[i-1] >= 'A' && pes[i-1] <= 'Z') || (pes[i-1] >= '0' && pes[i-1] <= '9'))) {
			if pes[i] >= 'a' && pes[i] <= 'z' {
				pes[i] = rune(pes[i] - 32)
			}
		} else if i == 0 {
			if pes[i] >= 'a' && pes[i] <= 'z' {
				pes[i] = rune(pes[i] - 32)
			}
		} else {
			if pes[i] >= 'A' && pes[i] <= 'Z' {
				pes[i] = rune(pes[i] + 32)
			}
		}
	}
	return string(pes)
}
