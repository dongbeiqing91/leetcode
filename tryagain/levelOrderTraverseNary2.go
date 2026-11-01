package tryagain

type State struct {
	node  *Node
	depth int
}

func levelNodeTraverseNaryState(root *Node) {
	if root == nil {
		return
	}
	q := []State{{node: root, depth: 1}}
	for len(q) > 0 {
		qLen := len(q)
		for range qLen {
			curState := q[0]
			q = q[1:]
			cur := curState.node
			for _, child := range cur.Children {
				if child != nil {
					q = append(q, State{
						node:  child,
						depth: curState.depth + 1,
					})
				}
			}
		}
	}
}
