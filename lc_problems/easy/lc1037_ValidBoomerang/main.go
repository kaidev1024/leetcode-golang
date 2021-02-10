func isBoomerang(points [][]int) bool {
	p0 := points[0]
	p1 := points[1]
	p2 := points[2]

	d0x := p0[0] - p1[0]
	d0y := p0[1] - p1[1]
	d1x := p0[0] - p2[0]
	d1y := p0[1] - p2[1]
	return d0x*d1y != d1x*d0y
}