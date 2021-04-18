/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := make(map[string][]*TreeNode)
	var helper func(root *TreeNode) string
	helper = func(root *TreeNode) string {
		if root == nil {
			return "e"
		}
		l := helper(root.Left)
		r := helper(root.Right)
		v := fmt.Sprintf("%dl%sr%s", root.Val, l, r)
		m[v] = append(m[v], root)
		return v
	}
	helper(root)

	ret := make([]*TreeNode, 0)
	for _, v := range m {
		n := len(v)
		if n == 1 {
			continue
		}
		flag := false
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if isSame(v[i], v[j]) {
					ret = append(ret, v[i])
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
	}
	return ret
}

func isSame(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}