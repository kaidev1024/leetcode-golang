func shoppingOffers(price []int, special [][]int, needs []int) int {
	nItems := len(price)

	var helper func(res []int, sum int) int
	helper = func(res []int, sum int) int {
		s := 0
		for i, v := range res {
			s += price[i] * v
		}
		if s == 0 {
			return sum
		}
		s += sum
		for _, sp := range special {
			newRes := make([]int, nItems)
			flag := false
			for i := 0; i < nItems; i++ {
				if res[i] < sp[i] {
					flag = true
					break
				}
				newRes[i] = res[i] - sp[i]
			}
			if flag {
				continue
			}
			s = min(s, helper(newRes, sum+sp[nItems]))
		}
		return s
	}

	return helper(needs, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}