package problems

func BetweenTwoSets(a, b []int32) int32 {
	var result int32
	var factors []int32

	// Find the factors of the smallest number in the second array
	for i := int32(1); i <= b[0]; i++ {
		if b[0]%i == 0 {
			factors = append(factors, i)
		}
	}

	// Filter out the factors that are not factors of all the numbers in the second array
	for i := 0; i < len(factors); i++ {
		for j := 0; j < len(b); j++ {
			if b[j]%factors[i] != 0 {
				factors = append(factors[:i], factors[i+1:]...)
				i--
				break
			}
		}
	}

	// Filter out the factors that are not multiples of all the numbers in the first array
	for i := 0; i < len(factors); i++ {
		for j := 0; j < len(a); j++ {
			if factors[i]%a[j] != 0 {
				factors = append(factors[:i], factors[i+1:]...)
				i--
				break
			}
		}
	}

	// Count the number of factors
	result = int32(len(factors))
	return result
}
