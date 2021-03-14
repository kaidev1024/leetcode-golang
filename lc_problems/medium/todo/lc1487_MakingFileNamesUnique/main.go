
func getFolderNames(names []string) []string {
	m := make(map[string]int)
	n := len(names)
	result := make([]string, n)
	for i, name := range names {
		if curV, ok := m[name]; ok {
			for {
				newName := fmt.Sprintf("%s(%v)", name, curV)
				if _, ok := m[newName]; ok {
					curV++
				} else {
					m[name] = curV
					m[newName] = 1
					result[i] = newName
					break
				}
			}
		} else {
			m[name] = 1
			result[i] = name
		}
	}
	return result
}