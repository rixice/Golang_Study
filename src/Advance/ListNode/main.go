package main

type ListNode struct {
	Val  int
	next *ListNode
}

func main() {

}

// 升序合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		head, tail, l1 = l1, l1, l1.next
	} else {
		head, tail, l2 = l2, l2, l2.next
	}

	// 循环，直到某一个链表已遍历完毕
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.next, l1 = l1, l1.next
		} else {
			tail.next, l2 = l2, l2.next
		}
		tail = tail.next
	}

	// 剩下的节点拼接到新链表尾部
	if l1 != nil {
		tail.next = l1
	} else if l2 != nil {
		tail.next = l2
	}

	return head
}
