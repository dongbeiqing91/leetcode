package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	first := ret
	increase := false
	for l1 != nil || l2 != nil {
		var added int
		if l1 == nil {
			added = l2.Val
		} else if l2 == nil {
			added = l1.Val
		} else {
			added = l1.Val + l2.Val
		}
		if increase {
			added += 1
		}
		if added >= 10 {
			increase = true
			added -= 10
		} else {
			increase = false
		}
		ret.Val = added
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil {
			ret.Next = new(ListNode)
			ret = ret.Next
		}
	}
	if increase {
		ret.Next = new(ListNode)
		ret.Next.Val = 1
	}
	return first
}
