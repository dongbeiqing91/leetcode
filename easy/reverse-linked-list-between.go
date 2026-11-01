package easy

// https://leetcode.cn/problems/reverse-linked-list-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	var prev = dummy
	for range left - 1 {
		prev = prev.Next
	}
	p1 := prev
	p2 := prev.Next
	cur := p2.Next
	prev = p2
	for range right - left {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	p1.Next = prev
	p2.Next = cur
	return dummy.Next
}
