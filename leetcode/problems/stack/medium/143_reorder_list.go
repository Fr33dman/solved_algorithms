package main

import "fmt"

/*
Reorder List
https://leetcode.com/problems/reorder-list/description/?envType=problem-list-v2&envId=stack

You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.



Example 1:


Input: head = [1,2,3,4]
Output: [1,4,2,3]
Example 2:


Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]


Constraints:

The number of nodes in the list is in the range [1, 5 * 104].
1 <= Node.val <= 1000
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	/*
		Time complexity: O(n)
		Space complexity: O(n)
	*/

	if head == nil || head.Next == nil {
		return
	}

	var start, cur *ListNode
	stack := make([]*ListNode, 0)
	cur = head
	for {
		if cur == nil {
			break
		}
		stack = append(stack, cur)
		cur = cur.Next
	}

	stack2 := stack[len(stack)/2:]
	stack = stack[:len(stack)/2]

	start = stack[0]
	cur = start
	cur.Next = stack2[len(stack2)-1]
	cur = cur.Next

	for i := 1; i < len(stack) || i < len(stack2); i++ {
		if i < len(stack) {
			cur.Next = stack[i]
			cur = cur.Next
		}
		if i < len(stack2) {
			cur.Next = stack2[len(stack2)-i-1]
			cur = cur.Next
		}
	}

	cur.Next = nil
}

func testReorderList() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	for {
		if head == nil {
			break
		}
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}
