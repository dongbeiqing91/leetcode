package easy

// 160. 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 时间 O(m+n)，额外空间 O(1)。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA, lengthB := 0, 0
	for node := headA; node != nil; node = node.Next {
		lengthA++
	}
	for node := headB; node != nil; node = node.Next {
		lengthB++
	}

	// 让较长的链表先走长度差，使两个指针到链表末尾的距离相同。
	for lengthA > lengthB {
		headA = headA.Next
		lengthA--
	}
	for lengthB > lengthA {
		headB = headB.Next
		lengthB--
	}

	// 首次指向同一个节点即为交点。无交点时，两者同时到达 nil。
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
