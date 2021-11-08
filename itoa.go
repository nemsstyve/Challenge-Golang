package piscine

func Itoa(n int) string {
	var str string
	var newstr string
	if n < 0 {
		n = n * -1
		newstr = "-"
		for i := n; i > 0; i /= 10 {
			r := rune(i%10 + 48)
			str += string(r)
		}
		for i := len(str) - 1; i >= 0; i-- {
			newstr += string(str[i])
		}
	} else if n == 0 {
		newstr = "0"
	} else if n > 0 {
		for i := n; i > 0; i /= 10 {
			r := rune(i%10 + 48)
			str += string(r)
		}
		for i := len(str) - 1; i >= 0; i-- {
			newstr += string(str[i])
		}
	}
	return newstr
}
