package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	for i := 0; i < len(a)-1 && ascending; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascending = false
		}
	}

	if !ascending {
		for i := 0; i < len(a)-1 && descending; i++ {
			if f(a[i], a[i+1]) < 0 {
				descending = false
			}
		}
	}
	return ascending || descending
}

/*func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	for i := 1; i < len(a) && ascending; i++ {
		if f(a[i-1], a[i]) > 0 {
			ascending = false
		}
	}

	if !ascending {
		for i := 1; i < len(a) && descending; i++ {
			if f(a[i-1], a[i]) < 0 {
				descending = false
			}
		}
	}
	return ascending || descending
}*/
