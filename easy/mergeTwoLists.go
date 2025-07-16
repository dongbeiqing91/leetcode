package easy

// https://leetcode.cn/problems/merge-two-sorted-lists/submissions/644554104/
// 21. 合并两个有序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	newList := &ListNode{}
	first := newList
	for list1 != nil || list2 != nil {
		newList.Next = &ListNode{}
		newList = newList.Next
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				newList.Val = list1.Val
				list1 = list1.Next
				continue
			} else {
				newList.Val = list2.Val
				list2 = list2.Next
				continue
			}
		} else {
			if list1 == nil {
				newList.Val = list2.Val
				list2 = list2.Next
				continue
			} else {
				newList.Val = list1.Val
				list1 = list1.Next
				continue
			}
		}
	}

	return first.Next
}
