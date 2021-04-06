func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		dict := make(map[int]int)
		for j := 0; j < len(points); j++ {
			d := dis(points[i], points[j])
			if v, ok := dict[d]; !ok {
				dict[d]++
			} else {
				res += 2 * v
				dict[d]++
			}
		}
	}
	return res
}

func dis(a []int, b []int) int {
	x := a[0] - b[0]
	y := a[1] - b[1]
	return x*x + y*y
}