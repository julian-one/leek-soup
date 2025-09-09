package easy

import "github.com/julian-one/leek-soup/problems"

// 104. Maximum Depth of a Binary Tree
// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
// Time Complexity: O(n) - visit each node exactly once
// Space Complexity: O(h) - where h is the height of the tree (recursion stack)
func MaxDepth(root *problems.TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}
