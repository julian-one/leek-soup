package easy

import "github.com/julian-one/leek-soup/problems"

// 94. Binary Tree Inorder Traversal (DFS)
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
//
// Time Complexity: O(n) - visit each node exactly once
// Space Complexity: O(h) - where h is the height of the tree (recursion stack)
func InorderTraversal(root *problems.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	result = append(result, InorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, InorderTraversal(root.Right)...)

	return result
}
