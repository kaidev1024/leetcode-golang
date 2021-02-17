
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)

	var helper func(root *TreeNode, target int, arr []int)
	helper = func(root *TreeNode, target int, arr []int) {
		target -= root.Val
		arr = append(arr, root.Val)
		if target == 0 && root.Left == nil && root.Right == nil {
			arrCopy := make([]int, len(arr))
			copy(arrCopy, arr)
			result = append(result, arrCopy)
			return
		}
		if root.Left != nil {
			helper(root.Left, target, arr)
		}
		if root.Right != nil {
			helper(root.Right, target, arr)
		}
	}
	helper(root, targetSum, []int{})
	return result
}