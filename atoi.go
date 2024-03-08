package piscine

func Atoi(s string) int {
	sign := 1

	out := 0
	for i, ch := range s {
		if ch >= '0' && ch <= '9' {
			num := ch - '0'
			out = out*10 + int(num)
		} else if ch == '-' && i == 0 {
			sign = -1
		} else if ch == '+' && i == 0 {
			sign = 1
		} else {
			return 0
		}
	}
	return out * sign
}
