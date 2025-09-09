package easy

import "github.com/julian-one/leek-soup/problems"

// 100. Same Tree (DFS)
// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
//
// Time Complexity: O(min(m,n)) - where m and n are the number of nodes in each tree
// Space Complexity: O(min(m,n)) - recursion stack depth
func IsSameTree(p *problems.TreeNode, q *problems.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
