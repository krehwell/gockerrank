package problems

func sumArray(arr []int32) int32 {
    var sum int32 = 0

    for i := range arr {
        sum += arr[i]
    }

    return sum
}
