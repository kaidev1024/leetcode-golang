func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.SliceStable(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	result := 0
	for i := 0; i < len(boxTypes); i++ {
		if truckSize > boxTypes[i][0] {
			result += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			result += boxTypes[i][1] * truckSize

			return result
		}
	}
	return result
}