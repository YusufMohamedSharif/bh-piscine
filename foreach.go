package piscine

func PrintNumber(a int) {
	a = a * 2
}

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a)-1; i++ {
		f(a[i])
	}
}
