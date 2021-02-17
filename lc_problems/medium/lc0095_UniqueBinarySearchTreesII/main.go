
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func add(root *TreeNode, v int) {
	if root == nil {
		return
	}
	root.Val += v
	add(root.Left, v)
	add(root.Right, v)
}

func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	result := &TreeNode{root.Val, nil, nil}
	result.Left = copyTree(root.Left)
	result.Right = copyTree(root.Right)
	return result
}

func generateTrees(n int) []*TreeNode {
	arr := make([][]*TreeNode, n+1)
	arr[0] = []*TreeNode{nil}
	arr[1] = []*TreeNode{&TreeNode{1, nil, nil}}

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			lArr := arr[j-1]
			rArr := arr[i-j]
			for _, l := range lArr {
				for _, r := range rArr {
					root := &TreeNode{j, nil, nil}
					root.Left = l
					root.Right = r
					newTree := copyTree(root)
					add(newTree.Right, j)
					arr[i] = append(arr[i], newTree)
				}
			}
		}
	}
	return arr[n]
}