package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	res := 1
	res2 := 1
	res3 := 1
	for k, v := range a {
		if k != len(a)-1 {
			if f(v, a[k+1]) < 0 {
				res++
			}
			if f(v, a[k+1]) > 0 {
				res2++
			}
			if f(v, a[k+1]) == 0 {
				res3++
			}
		}
	}
	return res == len(a) || res2 == len(a) || res3 == len(a)
}
