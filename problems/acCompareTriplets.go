package problems

func compareTriplets(a []int32, b []int32) []int32 {
    var aScore int32 = 0
    var bScore int32 = 0

    for i := range a {
        if a[i] > b[i] {
            aScore++;
        } else if b[i] > a[i] {
            bScore++;
        }
    }

    return []int32{aScore, bScore}
}
