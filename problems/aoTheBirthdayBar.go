package problems

import "fmt"

func sumSegment(a []int32) (sum int32) {
	fmt.Println(a)
	for _, v := range a {
		sum += v
	}

	return
}

// Two children, Lily and Ron, want to share a chocolate bar. Each of the squares has an integer on it.
// Lily decides to share a contiguous segment of the bar selected such that:
// - The length of the segment matches Ron's birth month, and,
// - The sum of the integers on the squares is equal to his birth day.
// Determine how many ways she can divide the chocolate.
func theBirthdayBarProblem(s []int32, d int32, m int32) int32 {
	p0 := int32(0)
	p1 := m

	count := int32(0)

	for ; p1 <= int32(len(s)); p0, p1 = p0+1, p1+1 {
		sum := sumSegment(s[p0:p1])

		if sum == d {
			count++
		}
	}

	return count
}
