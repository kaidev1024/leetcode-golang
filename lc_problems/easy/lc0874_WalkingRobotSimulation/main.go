func robotSim(commands []int, obstacles [][]int) int {
	obs := make(map[[2]int]struct{})
	for _, v := range obstacles {
		obs[[2]int{v[0], v[1]}] = struct{}{}
	}

	x, y := 0, 0
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	cur := 0
	res := 0
	for _, v := range commands {
		if v == -2 {
			cur = (cur + 3) % 4
		} else if v == -1 {
			cur = (cur + 1) % 4
		} else {
			for i := 0; i < v; i++ {
				xx := x + dir[cur][0]
				yy := y + dir[cur][1]
				if _, ok := obs[[2]int{xx, yy}]; ok {
					break
				} else {
					x, y = xx, yy
				}
			}
			if x*x+y*y > res {
				res = x*x + y*y
			}
		}
	}
	return res
}