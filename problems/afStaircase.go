package problems

// Its base and height are both equal to
// It is drawn using # symbols and spaces. The last line is not preceded by any spaces.
// Write a program that prints a staircase of size
func staircase(n int) string {
    result := ""

    currToHashIdx := n - 1;

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if j >= currToHashIdx {
                result += "#"
            } else {
                result += " "
            }
        }
        currToHashIdx--;
        result += "\n"
    }

    return result
}
