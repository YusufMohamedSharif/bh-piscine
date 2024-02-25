package piscine

func Sqrt(n int) int {
	if n < 0 {
		return 0
	}

	root := 0
	for i := 0; i*i <= n; i++ {
		if i*i == n {
			root = i
			break
		}
	}

	return root
}
