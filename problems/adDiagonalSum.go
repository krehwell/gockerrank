package problems

import "fmt"

func diagonalSum(arr [][]int32) int32 {
    // 00
    // 11
    // 22
    rSum := 0
    for i := 0; i < len(arr); i ++ {
        for j := 0; j < len(arr[i]); j++ {
            if i == j {
                rSum += int(arr[i][j])
            }
        }
    }

    // 02
    // 11
    // 20
    lSum := 0
    for i := 0; i < len(arr); i ++ {
        for j := 0; j < len(arr[i]); j++ {
            if i + j == len(arr[i]) - 1 {
                lSum += int(arr[i][j])
            }
        }
    }

    fmt.Println(rSum)
    fmt.Println(lSum)

    return 0
}
