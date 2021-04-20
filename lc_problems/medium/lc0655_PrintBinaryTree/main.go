/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	h := depth(root)
	n := 1<<h - 1
	res := make([][]string, h)
	for i := 0; i < len(res); i++ {
		res[i] = make([]string, n)
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = ""
		}
	}

	helper(root, 0, h, 0, n-1, res)
	return res
}

func helper(root *TreeNode, cur int, h int, start int, end int, res [][]string) {
	if cur == h {
		return
	}
	if root == nil {
		return
	}
	mid := (start + end) / 2
	res[cur][mid] = strconv.Itoa(root.Val)
	helper(root.Left, cur+1, h, start, mid-1, res)
	helper(root.Right, cur+1, h, mid+1, end, res)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}