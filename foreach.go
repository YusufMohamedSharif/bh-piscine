package piscine

func PrintNumber(a int) {
	a = a
}

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
}
