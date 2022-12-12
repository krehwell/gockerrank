package problems

// You are in charge of the cake for a child's birthday. You have decided the cake will
// have one candle for each year of their total age. They will only be able to
// blow out the tallest of the candles. Count how many candles are tallest.
func birthdayCandle(arr []int) int {
    max := -1
    numOfMax := 0

    for i := 0; i < len(arr); i++ {
        currHeight := arr[i]

        if currHeight > max {
            max = currHeight
            numOfMax = 1
        } else if currHeight == max {
            numOfMax++
        }
    }

    return numOfMax
}
