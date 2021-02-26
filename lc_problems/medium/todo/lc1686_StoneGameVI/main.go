func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	sum := 0
	for i, v := range aliceValues {
		aliceValues[i] = v + bobValues[i]
		sum += bobValues[i]
	}

	sort.Slice(aliceValues, func(i, j int) bool {
		return aliceValues[i] > aliceValues[j]
	})

	for i := 0; i < n; i += 2 {
		sum -= aliceValues[i]
	}
	if sum < 0 {
		return 1
	} else if sum > 0 {
		return -1
	}
	return 0
}