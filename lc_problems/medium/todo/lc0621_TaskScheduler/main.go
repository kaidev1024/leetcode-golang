func leastInterval(tasks []byte, n int) int {
	frequencies := make([]int, 26)
	for _, task := range tasks {
		frequencies[task-'A']++
	}

	MIN := func(v1, v2 int) int {
		if v1 < v2 {
			return v1
		}
		return v2
	}
	MAX := func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		}
		return v2
	}

	sort.Ints(frequencies)

	f_max := frequencies[25]
	idle_time := (f_max - 1) * n

	for i := len(frequencies) - 2; i >= 0 && idle_time > 0; i-- {
		idle_time -= MIN(frequencies[i], f_max-1)
	}
	idle_time = MAX(0, idle_time)

	return idle_time + len(tasks)
}