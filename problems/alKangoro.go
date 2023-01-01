package problems

// import "fmt"

func KangoroProblem(x1 int32, v1 int32, x2 int32, v2 int32) string {
	var min *int32 = nil
	var max *int32 = nil

	if x1 < x2 {
		min = &x1
		max = &x2
	} else {
		min = &x2
		max = &x1
	}

	for *min < *max {
		x1 += v1
		x2 += v2

		if x1 == x2 {
			return "YES"
		}
	}

	return "NO"
}
