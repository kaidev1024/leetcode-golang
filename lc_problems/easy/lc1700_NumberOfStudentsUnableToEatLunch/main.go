func countStudents(students []int, sandwiches []int) int {
	n := len(students)
	result := 0

	nMismatch := 0
	limit := n

	i := 0
	j := 0

	for ; j < n; i++ {
		if students[i] == sandwiches[j] {
			result++
			nMismatch = 0
			limit--
			j++
			continue
		}
		students = append(students, students[i])
		nMismatch++
		if limit == nMismatch {
			return n - result
		}
	}
	return n - result
}