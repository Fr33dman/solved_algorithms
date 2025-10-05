package main

import "fmt"

/*
Same Tree
https://leetcode.com/problems/same-tree/description/?envType=problem-list-v2&envId=tree

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.


Example 1:


Input: p = [1,2,3], q = [1,2,3]
Output: true
Example 2:


Input: p = [1,2], q = [1,null,2]
Output: false
Example 3:


Input: p = [1,2,1], q = [1,1,2]
Output: false


Constraints:

The number of nodes in both trees is in the range [0, 100].
-104 <= Node.val <= 104
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	/*
		Time complexity: O(n)
		Space complexity: O(n)
	*/

	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	stack1, stack2 := []*TreeNode{p}, []*TreeNode{q}
	var cursor1, cursor2 *TreeNode

	for len(stack1) != 0 && len(stack2) != 0 {
		cursor1 = stack1[len(stack1)-1]
		cursor2 = stack2[len(stack2)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]

		if cursor1.Val != cursor2.Val {
			return false
		}

		var v1, v2 *TreeNode
		for i := 0; i < 2; i++ {
			switch i {
			case 0:
				v1 = cursor1.Left
				v2 = cursor2.Left
			case 1:
				v1 = cursor1.Right
				v2 = cursor2.Right
			}

			x := v1 == nil
			y := v2 == nil

			if x && y {
				continue
			} else if (x || y) || v1.Val != v2.Val {
				return false
			} else {
				stack1 = append(stack1, v1)
				stack2 = append(stack2, v2)
			}
		}
	}
	return true
}

func testIsSameTree() {
	fmt.Println("test1: ", isSameTree(&TreeNode{Val: 0, Left: nil, Right: nil}, &TreeNode{Val: 0, Left: nil, Right: nil}) == true)
	fmt.Println("test2: ", isSameTree(&TreeNode{Val: 1, Left: nil, Right: nil}, &TreeNode{Val: 0, Left: nil, Right: nil}) == false)
	fmt.Println("test3: ", isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}, Right: nil},
	) == false)
	fmt.Println("test3: ", isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
	) == true)
}
