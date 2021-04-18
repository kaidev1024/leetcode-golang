func exclusiveTime(n int, logs []string) []int {
	times := make([]int, n)

	stack := []int{}
	prevTime := 0

	for _, log := range logs {
		parts := strings.Split(log, ":")

		function, _ := strconv.Atoi(parts[0])
		eventType := parts[1]
		timestamp, _ := strconv.Atoi(parts[2])

		if eventType == "start" {
			if len(stack) != 0 {
				times[stack[len(stack)-1]] += timestamp - prevTime
			}

			stack = append(stack, function)
			prevTime = timestamp
		} else {
			times[stack[len(stack)-1]] += timestamp - prevTime + 1
			prevTime = timestamp + 1
			stack = stack[:len(stack)-1]
		}
	}

	return times
}