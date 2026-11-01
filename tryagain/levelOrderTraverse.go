package tryagain

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://labuladong.online/zh/algo/data-structure-basic/binary-tree-traverse-basic

func levelOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	var depth int
	for len(q) > 0 {
		queueLen := len(q)
		for i := 0; i < queueLen; i++ {
			cur := q[0]
			fmt.Printf("current node: %v, level: %d\n", cur, depth)
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
}
