type Solution struct {
	rects [][]int
	areas []int
}

func Constructor(rects [][]int) Solution {
	areas := []int{0}
	preArea := 0
	for _, rect := range rects {
		area := getArea(rect)
		fmt.Println(area)
		preArea += area
		areas = append(areas, preArea)
	}
	return Solution{
		rects: rects,
		areas: areas,
	}
}

func (this *Solution) Pick() []int {
	areas := this.areas
	total := areas[len(areas)-1]
	randNum := rand.Intn(total) + 1
	index := sort.SearchInts(areas, randNum) - 1
	k := randNum - areas[index] - 1
	rect := this.rects[index]
	width := dif(rect[2], rect[0]) + 1
	row := k / width
	col := k % width
	return []int{col + min(rect[0], rect[2]), row + min(rect[1], rect[3])}
}

func dif(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getArea(rect []int) int {

	return (dif(rect[2], rect[0]) + 1) * (dif(rect[1], rect[3]) + 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */