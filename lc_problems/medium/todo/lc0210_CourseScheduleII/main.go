func findOrder(numCourses int, prerequisites [][]int) []int {
	indegree := make([]int, numCourses)
	outdegree := make([][]int, numCourses)
	result := make([]int, 0, numCourses)

	for _, pair := range prerequisites {
		indegree[pair[0]]++
		outdegree[pair[1]] = append(outdegree[pair[1]], pair[0])
	}

	count := 0

	n := 0
	root := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			n++
			root = append(root, i)
		}
	}

	for n > 0 {
		count += n
		result = append(result, root...)
		n = 0
		newRoot := make([]int, 0)
		for _, v := range root {
			for _, index := range outdegree[v] {
				indegree[index]--
				if indegree[index] == 0 {
					n++
					newRoot = append(newRoot, index)
				}
			}
		}
		root = newRoot
	}
	if count == numCourses {
		return result
	}
	return nil
}