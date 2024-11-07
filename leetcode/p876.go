package leetcode

func MiddeNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func MiddleNode2(head *ListNode) *ListNode {
	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	return arr[len(arr)/2]
}
