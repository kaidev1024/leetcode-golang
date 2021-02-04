func reformatDate(date string) string {
	fields := strings.Fields(date)
	day := fields[0][0] - '0'
	if fields[0][1] <= '9' {
		day = day*10 + fields[0][1] - '0'
	}

	m := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	return fmt.Sprintf("%s-%s-%02d", fields[2], m[fields[1]], day)
}