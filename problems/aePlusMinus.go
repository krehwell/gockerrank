package problems

import (
	"fmt"
	"strconv"
)

func plusMinus(arr []int32) {
	ratio := len(arr)

	var numOfPos float64 = 0
	var numOfNeg float64 = 0
	var numOfZero float64 = 0

	for i := range arr {
		if arr[i] == 0 {
			numOfZero++
		} else if arr[i] > 0 {
			numOfPos++
		} else if arr[i] < 0 {
			numOfNeg++
		}
	}

	posRatio := numOfPos / float64(ratio)
	negRatio := numOfNeg / float64(ratio)
	zeroRatio := numOfZero / float64(ratio)

	a := strconv.FormatFloat(posRatio, 'f', ratio, 64)
	b := strconv.FormatFloat(negRatio, 'f', ratio, 64)
	c := strconv.FormatFloat(zeroRatio, 'f', ratio, 64)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
