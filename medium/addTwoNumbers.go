package medium

// https://leetcode.cn/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 在两条链表上的指针
	p1, p2 := l1, l2
	// 虚拟头结点（构建新链表时的常用技巧）
	dummy := &ListNode{-1, nil}
	// 指针 p 负责构建新链表
	p := dummy
	// 记录进位
	carry := 0
	// 开始执行加法，两条链表走完且没有进位时才能结束循环
	for p1 != nil || p2 != nil || carry > 0 {
		// 先加上上次的进位
		val := carry
		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}
		// 处理进位情况
		carry = val / 10
		val = val % 10
		// 构建新节点
		p.Next = &ListNode{val, nil}
		p = p.Next
	}
	// 返回结果链表的头结点（去除虚拟头结点）
	return dummy.Next
}
