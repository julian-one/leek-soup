package easy

import "github.com/julian-one/leek-soup/problems"

// 111. Minimum Depth of a Binary Tree (BFS)
// Given a binary tree, find its minimum depth.
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
// Note: A leaf is a node with no children.
//
// Time Complexity: O(n) - in worst case, visit all nodes
// Space Complexity: O(w) - where w is the maximum width of the tree (queue size)
func MinDepth(root *problems.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*problems.TreeNode{root}
	depth := 1

	for len(queue) > 0 {
		level := len(queue)

		for range level {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
