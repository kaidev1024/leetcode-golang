func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	n := len(nodes)

	stack := make([]string, 0, n)
	pos := 0
	for i := n - 1; i >= 0; i-- {
		if nodes[i] == "#" {
			stack = append(stack, "#")
			pos++
			continue
		}
		if pos < 2 || stack[pos-1] != "#" || stack[pos-2] != "#" {
			return false
		}
		pos--
		stack = stack[:pos]
	}
	return len(stack) == 1 && stack[0] == "#"
}   