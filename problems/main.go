package problems

import "fmt"

func Solve() {
	result := diagonalSum([][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	})
	fmt.Println(result)
}