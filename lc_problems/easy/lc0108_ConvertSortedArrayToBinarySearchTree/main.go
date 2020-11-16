import (
	"fmt"

	"github.com/kaidev1024/leetcode-golang/structs/tree"
)

func util(nums []int, left, right int) *tree.TreeNode {

	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	node := &TreeNode{
		Val:   nums[mid],
		Left:  util(nums, left, mid-1),
		Right: util(nums, mid+1, right),
	}

	return node
}

func sortedArrayToBST(nums []int) *tree.TreeNode {
	return util(nums, 0, len(nums)-1)
}