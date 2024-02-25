package piscine

func FindNextPrime(nb int) int {
	isPrime := func(num int) bool {
		if num <= 1 {
			return false
		}
		if num <= 3 {
			return true
		}
		if num%2 == 0 || num%3 == 0 {
			return false
		}
		i := 5
		for i*i <= num {
			if num%i == 0 || num%(i+2) == 0 {
				return false
			}
			i += 6
		}
		return true
	}

	if nb <= 1 {
		return 2
	}

	prime := nb
	found := false
	for !found {
		prime++
		if isPrime(prime) {
			found = true
		}
	}
	return prime
}
