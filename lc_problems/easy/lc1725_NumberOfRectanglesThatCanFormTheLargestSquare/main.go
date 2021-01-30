func countGoodRectangles(rectangles [][]int) int {
	count := 0
	length := 0
	l := len(rectangles)
	for i := 0; i < l; i++ {
		rect := rectangles[i]
		curLength := rect[0]
		if curLength > rect[1] {
			curLength = rect[1]
		}

		if curLength == length {
			count++
		}
		if curLength > length {
			length = curLength
			count = 1
		}
	}
	return count
}