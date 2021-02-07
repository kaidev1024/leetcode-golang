type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	count := 0
	val := 0
	maxCount := -1
	result := make([]int, 0)

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root.Left != nil {
			inorder(root.Left)
		}
		if root.Val == val {
			count++
		} else {
			if count == maxCount {
				result = append(result, val)
				val = root.Val
				count = 1
			} else if count > maxCount {
				result = result[0:0]
				result = append(result, val)
				val = root.Val
				maxCount = count
				count = 1

			} else {
				val = root.Val
				count = 1
			}
		}
		if root.Right != nil {
			inorder(root.Right)
		}
	}
	inorder(root)
	if count == maxCount {
		result = append(result, val)
	} else if count > maxCount {
		return []int{val}
	}

	return result
}