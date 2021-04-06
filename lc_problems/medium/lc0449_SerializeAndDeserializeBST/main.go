/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	if root == nil {
		return ""
	}
	return strconv.Itoa(root.Val) + "(" + this.serialize(root.Left) + ")(" + this.serialize(root.Right) + ")"
}

// Deserializes your encoded data to tree.

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	pos := 0
	n := len(data)

	getInt := func() int {
		result := 0
		for pos < n && data[pos] >= '0' && data[pos] <= '9' {
			result *= 10
			result += int(data[pos] - '0')
			pos++
		}
		return result
	}

	var helper func() *TreeNode
	helper = func() *TreeNode {
		if data[pos] == ')' {
			pos++
			return nil
		}
		val := getInt()
		tn := &TreeNode{
			Val: val,
		}

		if pos == n {
			return tn
		}
		if data[pos] == ')' {
			pos++
			return nil
		}
		fmt.Println(pos)
		pos++
		tn.Left = helper()
		pos++
		tn.Right = helper()
		pos++
		return tn
	}
	return helper()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */