func readBinaryWatch(num int) []string {
	mapHour := map[int][]string{}
	mapMinute := map[int][]string{}

	mapHour[0] = append(mapHour[0], "0")
	mapMinute[0] = append(mapMinute[0], "00")

	for i := 1; i < 60; i++ {
		num := i
		count := 0
		for num != 0 {
			if num&0x1 != 0 {
				count++
			}
			num >>= 1
		}
		if i <= 11 {
			mapHour[count] = append(mapHour[count], strconv.Itoa(i))
		}
		if i < 10 {
			mapMinute[count] = append(mapMinute[count], "0"+strconv.Itoa(i))
		} else {
			mapMinute[count] = append(mapMinute[count], strconv.Itoa(i))
		}
	}

	timeString := []string{}
	for hour := max(0, num-5); hour <= min(3, num); hour++ {
		for _, hString := range mapHour[hour] {
			for _, mString := range mapMinute[num-hour] {
				timeString = append(timeString, hString+":"+mString)
			}
		}
	}
	return timeString
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}