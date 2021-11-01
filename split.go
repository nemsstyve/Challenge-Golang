package piscine

func Split(s, sep string) []string {
	ln := 0
	ok := true
	ln_s := 0
	for c := range s {
		ln_s = c + 1
	}
	for i := 0; i < ln_s; i++ {
		same := true
		ln2 := 0
		for j, t := range sep {
			if rune(s[i+j]) != t {
				same = false
				break
			}
			ln2 = j
		}
		if same {
			if !ok {
				ln++
			}
			i += (ln2)
			ok = true
			continue
		}
		ok = false
	}
	ans := make([]string, ln+1)
	x := 0
	for i := 0; i < ln_s; i++ {
		same := true
		ln2 := 0
		for j, t := range sep {
			if rune(s[i+j]) != t {
				same = false
				break
			}
			ln2 = j
		}
		if same {
			if !ok {
				x++
			}
			i += (ln2)
			ok = true
			continue
		}
		ans[x] = ans[x] + string(s[i])
		ok = false
	}
	return ans
}
