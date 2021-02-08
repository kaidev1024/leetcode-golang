func matrixReshape(nums [][]int, r int, c int) [][]int {
	rOld := len(nums)
	if rOld == 0 {
		return nil
	}
	cOld := len(nums[0])

	if rOld*cOld != r*c {
		return nums
	}

	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	for rOldPos, cOldPos, rNewPos, cNewPos := 0, 0, 0, 0; rOldPos < rOld; {
		result[rNewPos][cNewPos] = nums[rOldPos][cOldPos]
		cNewPos++
		if cNewPos == c {
			cNewPos = 0
			rNewPos++
		}
		cOldPos++
		if cOldPos == cOld {
			cOldPos = 0
			rOldPos++
		}
	}

	return result
}