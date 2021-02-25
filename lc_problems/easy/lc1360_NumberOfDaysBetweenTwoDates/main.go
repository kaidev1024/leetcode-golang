func isLeapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}

func daysBetweenDates(date1 string, date2 string) int {
	if date1 > date2 {
		date1, date2 = date2, date1
	}
	n1 := dayOfYear(date1)
	n2 := dayOfYear(date2)
	result := n2 - n1

	y1, _ := strconv.Atoi(date1[:4])
	y2, _ := strconv.Atoi(date2[:4])
	for i := y1; i < y2; i++ {
		if isLeapYear(i) {
			result += 1
		}
		result += 365
	}
	return result
}

func dayOfYear(date string) int {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, _ := strconv.Atoi(date[:4])
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		months[1] = 29
	}

	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])

	result := 0
	for i := 0; i < month-1; i++ {
		result += months[i]
	}
	return result + day
}