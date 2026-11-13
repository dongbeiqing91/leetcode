package medium

// Node 是带父指针的二叉树节点。
type Node struct {
	Val    int
	Parent *Node
}

// 1650. 二叉树的最近公共祖先 III
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-iii/
// 时间 O(h)，其中 h 为树高。额外空间 O(1)。
func lowestCommonAncestor(p, q *Node) *Node {
	// 计算两条父指针链的长度，包含起点和根节点。
	lengthP, lengthQ := 0, 0
	for node := p; node != nil; node = node.Parent {
		lengthP++
	}
	for node := q; node != nil; node = node.Parent {
		lengthQ++
	}

	// 让更深的节点先向上走到相同深度。
	for lengthP > lengthQ {
		p = p.Parent
		lengthP--
	}
	for lengthQ > lengthP {
		q = q.Parent
		lengthQ--
	}

	// 同步向上走，第一次相遇的节点就是最近公共祖先。
	for p != q {
		p = p.Parent
		q = q.Parent
	}
	return p
}
