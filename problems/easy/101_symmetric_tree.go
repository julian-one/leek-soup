package easy

import "github.com/julian-one/leek-soup/problems"

// 101. Symmetric Tree
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
//
// Time Complexity: O(n) - visit each node exactly once
// Space Complexity: O(h) - where h is the height of the tree (recursion stack)
func IsSymmetric(root *problems.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(l *problems.TreeNode, r *problems.TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val &&
		isMirror(l.Left, r.Right) &&
		isMirror(l.Right, r.Left)
}
