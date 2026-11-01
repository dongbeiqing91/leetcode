package easy

// https://leetcode.cn/problems/minimum-depth-of-binary-tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	p := []*TreeNode{root}
	depth := 1
	for len(p) > 0 {
		pLen := len(p)
		for range pLen {
			cur := p[0]
			p = p[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				p = append(p, cur.Left)
			}
			if cur.Right != nil {
				p = append(p, cur.Right)
			}
		}
		depth++
	}
	return depth
}
