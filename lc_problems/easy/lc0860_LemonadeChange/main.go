func lemonadeChange(bills []int) bool {
	c5 := 0
	c10 := 0
	for _, v := range bills {
		if v == 5 {
			c5++
			continue
		}
		if v == 10 {
			if c5 == 0 {
				return false
			}
			c5--
			c10++
			continue
		}

		if c5 == 0 {
			return false
		}
		if c10 > 0 {
			c10--
			c5--
			continue
		}
		if c5 < 3 {
			return false
		}
		c5 -= 3
	}
	return true
}