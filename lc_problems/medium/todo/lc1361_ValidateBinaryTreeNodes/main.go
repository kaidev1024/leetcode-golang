func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	nodes := make([]int, n)
	for i := 0; i < n; i++ {
		v := leftChild[i]
		if v != -1 {
			if nodes[v] > 0 {
				return false
			}
			nodes[v]++
		}
		v = rightChild[i]
		if v != -1 {
			if nodes[v] > 0 {
				return false
			}
			nodes[v]++
		}
	}
	root := -1
	for i := 0; i < n; i++ {
		if nodes[i] == 0 {
			if root == -1 {
				root = i
			} else {
				return false
			}
		}
	}

	count := 0
	isUsed := make(map[int]struct{})
	var dfs func(root int) bool
	dfs = func(root int) bool {
		if root == -1 {
			return true
		}
		if _, ok := isUsed[root]; ok {
			return false
		}

		count++
		isUsed[root] = struct{}{}
		return dfs(leftChild[root]) && dfs(rightChild[root])
	}
	return dfs(root) && count == n
}