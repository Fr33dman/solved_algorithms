package main

/*
Binary Tree Inorder Traversal
https://leetcode.com/problems/binary-tree-inorder-traversal/description/?envType=problem-list-v2&envId=tree

Given the root of a binary tree, return the inorder traversal of its nodes' values.


Example 1:

Input: root = [1,null,2,3]

Output: [1,3,2]

Explanation:


Example 2:

Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

Output: [4,2,6,5,7,1,3,9,8]

Explanation:


Example 3:

Input: root = []

Output: []

Example 4:

Input: root = [1]

Output: [1]



Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	/*
		Recursive solution
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	if root == nil {
		return nil
	}
	res := append(inorderTraversal(root.Left), root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func inorderTraversalV2(root *TreeNode) []int {
	/*
		Iterative solution
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	var res []int

	stack := make([]*TreeNode, 0)
	cursor := root

	for len(stack) > 0 || cursor != nil {
		for cursor != nil {
			stack = append(stack, cursor)
			cursor = cursor.Left
		}

		cursor = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, cursor.Val)

		cursor = cursor.Right
	}

	return res
}

// No local test sorry, creating testing data sets will take too much time
