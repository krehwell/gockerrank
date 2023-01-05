package problems

// Maria plays college basketball and wants to go pro.
// Each season she maintains a record of her play.
// She tabulates the number of times she breaks her season record for most points and least points in a game. Points scored in the first game establish her record for the season, and she begins counting from there.

func breakingRecords(scores []int32) []int32 {
	highest := scores[0]
	lowest := scores[0]

	numOfHigh := int32(0)
	numOfLow := int32(0)

	for _, v := range scores {
		if v > highest {
			highest = v
			numOfHigh++
		} else if v < lowest {
			lowest = v
			numOfLow++
		}
	}

	return []int32{numOfHigh, numOfLow}
}
