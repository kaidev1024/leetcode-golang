func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	nhouses := len(houses)
	nheaters := len(heaters)

	ret := 0
	if houses[0] < heaters[0] {
		ret = max(ret, heaters[0]-houses[0])
	}
	if houses[nhouses-1] > heaters[nheaters-1] {
		ret = max(ret, houses[nhouses-1]-heaters[nheaters-1])
	}
	start := 0
	for start < nhouses {
		if houses[start] <= heaters[0] {
			start++
		} else {
			break
		}
	}
	if start == nhouses {
		return ret
	}
	end := nhouses - 1
	for end >= start {
		if houses[end] >= heaters[nheaters-1] {
			end--
		} else {
			break
		}
	}
	if start > end {
		return ret
	}

	for i, j := start, 0; i <= end; i++ {
		for houses[i] > heaters[j+1] {
			j++
		}
		dist := min(heaters[j+1]-houses[i], houses[i]-heaters[j])
		ret = max(dist, ret)

	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}