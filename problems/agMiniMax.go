package problems

import (
	"math"
)

func sumOfAllArr(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}

func removeIdx(_arr []int, idx int) []int {
	arr := make([]int, len(_arr))
	copy(arr, _arr)

	result := append(arr[:idx], arr[idx+1:]...)

	return result
}

// Given five positive integers, find the minimum and maximum values that can be calculated by
// summing exactly four of the five integers.
// Then print the respective minimum and maximum values as a single line of two space-separated long integers.
func minimax(arr []int) [2]int {
	minResult := math.Inf(1)
	maxResult := 0

	for i := 0; i < len(arr); i++ {
		currSum := sumOfAllArr(removeIdx(arr, i))

		if currSum > maxResult {
			maxResult = currSum
		}

		if float64(currSum) < minResult {
			minResult = float64(currSum)
		}
	}

	return [2]int{int(minResult), int(maxResult)}
}
