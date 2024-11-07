package main

import (
	"dsa/dsa1/leetcode"
)

func main() {
	a := leetcode.ListNode{Val: 1}
	a.Next = &leetcode.ListNode{Val: 2}
	a.Next.Next = &leetcode.ListNode{Val: 3}
	a.Next.Next.Next = &leetcode.ListNode{Val: 4}
	a.Next.Next.Next.Next = &leetcode.ListNode{Val: 5}
	leetcode.ReverseList(&a)
}
