package piscine

func BasicAtoi2(s string) int {
	x := 0
	for _, z := range s {
		if z >= '0' && z <= '9' {
			y := int(z - '0')
			x = ((x * 10) + y)
		} else {
			return 0
		}
	}
	return x
}
