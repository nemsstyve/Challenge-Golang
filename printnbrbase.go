package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	ans := ""
	ln := 0
	for _, c := range base {
		if c == c {
			ln++
		}
	}
	mx_p := ln
	if nbr < 0 {
		ans = "-"
		mx_p *= -1
	}
	if ln > 1 {
		for nbr/mx_p >= ln {
			mx_p *= ln
		}
		for mx_p != 0 {
			ans = ans + string(base[nbr/mx_p])
			nbr = nbr - nbr/mx_p*mx_p
			mx_p /= ln
		}
		x := map[rune]bool{}
		for _, c := range base {
			if c == '+' || c == '-' {
				ans = "NV"
				break
			}
			if x[c] == false {
				x[c] = true
			} else {
				ans = "NV"
				break
			}
		}
	} else {
		ans = "NV"
	}
	for _, c := range ans {
		z01.PrintRune(c)
	}
}
