func canChoose(groups [][]int, nums []int) bool {
	ng := len(groups)
	gi := 0
	pos := 0
	n := len(nums)
	for pos < n && gi < ng {
		if nums[pos] != groups[gi][0] {
			pos++
			continue
		}
		ind := 0
		ni := len(groups[gi])
		posi := pos
		for posi < n && ind < ni {
			if nums[posi] == groups[gi][ind] {
				posi++
				ind++
				if ind == ni {
					gi++
					break

				}
			} else {
				break
			}
		}
		if ind == ni {
			pos = posi
		} else {
			pos++
		}
	}
	return gi == ng
}