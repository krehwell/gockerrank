package problems

func addVector(by int32, vector []int32) (result []int32) {
	for i := range vector {
		result = append(result, vector[i]+by)
	}

	return
}

func countNearHouse(min, max int32, fruits []int32) (count int32) {
	for _, dist := range fruits {
		if dist >= min && dist <= max {
			count++
		}
	}

	return
}

/**
 * https://www.hackerrank.com/challenges/apple-and-orange/problem?isFullScreen=true
 */
func appleAndOrange(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) []int32 {
	appleDistList := addVector(a, apples)
	orangeDistList := addVector(b, oranges)

	appleCount := countNearHouse(s, t, appleDistList)
	orangeCount := countNearHouse(s, t, orangeDistList)

	return []int32{orangeCount, appleCount}
}
