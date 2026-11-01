package tryagain

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

// https://labuladong.online/zh/algo/data-structure-basic/n-ary-tree-traverse-basic

func levelNodeTraverseNary(root *Node) {
	if root == nil {
		return
	}
	q := []*Node{root}
	depth := 0
	for len(q) > 0 {
		qLen := len(q)
		for range qLen {
			cur := q[0]
			fmt.Printf("current node: %v, level: %d\n", cur, depth)
			q = q[1:]
			for _, child := range cur.Children {
				if child != nil {
					q = append(q, child)
				}
			}
		}
		depth++
	}
}

