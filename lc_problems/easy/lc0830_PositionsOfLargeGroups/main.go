func largeGroupPositions(S string) [][]int {

	res := [][]int{}
	index := 0

	for len(S) > index {
		count := 0
		for _, V := range S[index+1:] {
			if S[index] == byte(V) {
				count++
			} else {
				break
			}
		}
		if count >= 2 {
			res = append(res, []int{index, index + count})
			index = index + count + 1
			continue
		}
		index++
	}
	return res
}