func constructRectangle(area int) []int {
	r := int(math.Floor((math.Sqrt(float64(area)))))
	for i := r; i >= 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return []int{area, 1}
}
