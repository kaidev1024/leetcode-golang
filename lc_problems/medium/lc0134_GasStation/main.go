func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	start := 0
	end := 1 % n

	res := gas[0] - cost[0]

	for start != end {
		if res <= 0 {
			start = (start + n - 1) % n
			res += gas[start] - cost[start]
		} else {
			res += gas[end] - cost[end]
			end = (end + 1) % n
		}
	}
	if res < 0 {
		return -1
	}
	return start
}