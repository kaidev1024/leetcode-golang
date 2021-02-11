
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	m := make(map[int]*Employee)
	for i, n := 0, len(employees); i < n; i++ {
		m[employees[i].Id] = employees[i]
	}

	result := 0

	var dfs func(id int)

	dfs = func(id int) {
		e := m[id]
		result += e.Importance
		for _, v := range e.Subordinates {
			dfs(v)
		}
	}

	dfs(id)
	return result
}