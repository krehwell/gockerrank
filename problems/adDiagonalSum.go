package problems

import (
	"math"
)

func diagonalSum(arr [][]int32) int32 {
    aDiagonalSum := 0
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr[i]); j++ {
            if i == j {
                aDiagonalSum += int(arr[i][j])
            }
        }
    }

    bDiagonalSum := 0
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr[i]); j++ {
            if i + j == len(arr[i]) - 1 {
                bDiagonalSum += int(arr[i][j])
            }
        }
    }

    return int32(math.Abs(float64(aDiagonalSum) - float64(bDiagonalSum)))
}
