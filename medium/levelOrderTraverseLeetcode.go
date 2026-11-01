package medium

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-level-order-traversal/description

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var list [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		var levelList []int
		qLen := len(q)
		for range qLen {
			cur := q[0]
			q = q[1:]
			levelList = append(levelList, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		list = append(list, levelList)
	}
	return list
}
