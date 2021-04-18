
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	len := len(nums)
	l, r := 0, len-1
	return doConstructor(nums, l, r)
}

func doConstructor(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	i := GetMaxValueIndex(nums, l, r)
	node := &TreeNode{Val: nums[i]}
	node.Left = doConstructor(nums, l, i-1)
	node.Right = doConstructor(nums, i+1, r)
	return node
}

func GetMaxValueIndex(nums []int, l, r int) int { // [l..r]
	mvi := l
	for i := l; i <= r; i++ {
		if nums[i] > nums[mvi] {
			mvi = i
		}
	}
	return mvi
}