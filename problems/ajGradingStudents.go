package problems

func getListMulOf5() []int {
	bucket := []int{}

	for i := 0; i < 100; i = i + 5 {
		bucket = append(bucket, i+5)
	}

	return bucket
}

func findNearestOfList(n int, list []int) int {
	for i := 0; i < len(list); i++ {
		if list[i] > n {
			return list[i]
		}
	}

	return list[0]
}

func getFinalGrade(grade int) int {
	nearest := findNearestOfList(grade, getListMulOf5())

    if grade > 100 {
        return grade
    } else if grade < 38 {
		return grade
	} else if nearest-grade < 3 {
		return nearest
	} else {
		return grade
	}
}

// Sam is a professor at the university and likes to round each student's
// according to these rules:
// If the difference between the
// and the next multiple of is less than , round up to the next multiple of .
// If the value of is less than , no rounding occurs as the result will still be a failing grade.
func gradingStudents(grades []int) []int {
    results := []int{}
    for i := 0; i < len(grades); i++ {
        results = append(results, getFinalGrade(grades[i]))
    }

    return results
}
