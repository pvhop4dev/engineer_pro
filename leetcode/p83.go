package leetcode

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func DeleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	fast := head.Next
	for fast != nil {
		if current.Val == fast.Val {
			current.Next = fast.Next
		} else {
			current = fast
		}
		fast = fast.Next
	}
	return head
}
