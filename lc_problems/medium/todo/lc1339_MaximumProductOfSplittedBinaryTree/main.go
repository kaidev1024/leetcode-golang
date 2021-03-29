type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	var total, product int
	var nodeProduct func(root *TreeNode) int
	nodeProduct = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := root.Val + nodeProduct(root.Left) + nodeProduct(root.Right)
		product = max(product, (total-sum)*sum)
		return sum
	}

	var nodeSum func(root *TreeNode) int
	nodeSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := root.Val + nodeSum(root.Left) + nodeSum(root.Right)
		return sum
	}

	total = nodeSum(root)
	nodeProduct(root)

	return product % (int(1e9) + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}