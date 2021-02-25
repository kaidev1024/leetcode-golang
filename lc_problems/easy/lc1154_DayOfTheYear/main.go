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