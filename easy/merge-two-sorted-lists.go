package easy

// https://leetcode.cn/problems/merge-two-sorted-lists/description/
// 21. 合并两个有序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
    head := dummy
    for {
        if list1 == nil {
            head.Next = list2
            break
        }
        if list2 == nil {
            head.Next = list1
            break
        }
        if list1.Val <= list2.Val {
            head.Next = list1
            list1 = list1.Next
        } else {
            head.Next = list2
            list2 = list2.Next
        }
        head = head.Next
    }
    return dummy.Next
}
