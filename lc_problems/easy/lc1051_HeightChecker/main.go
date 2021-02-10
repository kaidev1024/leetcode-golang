func heightChecker(heights []int) int {
	moveCounter := 0
	copyHeights := make([]int, len(heights))
	copy(copyHeights, heights)
	sort.Ints(copyHeights)
	for i, v := range copyHeights {
		if v != heights[i] {
			moveCounter++
		}
	}
	return moveCounter
}