package piscine

func AtoiBase(s string, base string) int {
	ans := 0
	ln := 0
	st := map[rune]bool{}
	for _, c := range base {
		if st[c] || c == '-' || c == '+' {
			return ans
		}
		st[c] = true
		ln++
	}
	if ln > 1 {
		for _, c := range s {
			cnt := 0
			if st[c] {
				for _, j := range base {
					if j == c {
						break
					}
					cnt++
				}
				ans = ans*ln + cnt
			}
		}
	}
	return ans
}
