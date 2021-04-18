func predictPartyVictory(senater string) string {
	r := 0
	d := 0
	count := 0
	build := []byte{}
	senate := []byte(senater)

	for {
		for idx := 0; idx < len(senate); idx += 1 {
			if senate[idx] == 'D' {
				count -= 1
				if count < 0 {
					build = append(build, senate[idx])
					d += 1
				}
			} else {
				count += 1
				if count > 0 {
					r += 1
					build = append(build, senate[idx])
				}
			}
		}
		// fmt.Println(build)
		if r == 0 {
			return "Dire"
		}
		if d == 0 {
			return "Radiant"
		}
		senate, build = build, []byte{}
		r, d = 0, 0
	}
	return ""
}