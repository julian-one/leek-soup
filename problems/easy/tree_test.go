package easy

import (
	"reflect"
	"testing"

	"github.com/julian-one/leek-soup/problems"
)

// Helper function to create a tree from array representation
func createTree(vals []interface{}) *problems.TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &problems.TreeNode{Val: vals[0].(int)}
	queue := []*problems.TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) && vals[i] != nil {
			node.Left = &problems.TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(vals) && vals[i] != nil {
			node.Right = &problems.TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		tree []interface{}
		want []int
	}{
		{
			name: "Example 1",
			tree: []interface{}{1, nil, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "Empty tree",
			tree: []interface{}{},
			want: []int{},
		},
		{
			name: "Single node",
			tree: []interface{}{1},
			want: []int{1},
		},
		{
			name: "Complete tree",
			tree: []interface{}{1, 2, 3, 4, 5, 6, 7},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.tree)
			if got := InorderTraversal(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    []interface{}
		q    []interface{}
		want bool
	}{
		{
			name: "Example 1 - same trees",
			p:    []interface{}{1, 2, 3},
			q:    []interface{}{1, 2, 3},
			want: true,
		},
		{
			name: "Example 2 - different structure",
			p:    []interface{}{1, 2},
			q:    []interface{}{1, nil, 2},
			want: false,
		},
		{
			name: "Example 3 - different values",
			p:    []interface{}{1, 2, 1},
			q:    []interface{}{1, 1, 2},
			want: false,
		},
		{
			name: "Both empty",
			p:    []interface{}{},
			q:    []interface{}{},
			want: true,
		},
		{
			name: "One empty",
			p:    []interface{}{1},
			q:    []interface{}{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pTree := createTree(tt.p)
			qTree := createTree(tt.q)
			if got := IsSameTree(pTree, qTree); got != tt.want {
				t.Errorf("IsSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name string
		tree []interface{}
		want bool
	}{
		{
			name: "Example 1 - symmetric",
			tree: []interface{}{1, 2, 2, 3, 4, 4, 3},
			want: true,
		},
		{
			name: "Example 2 - not symmetric",
			tree: []interface{}{1, 2, 2, nil, 3, nil, 3},
			want: false,
		},
		{
			name: "Single node",
			tree: []interface{}{1},
			want: true,
		},
		{
			name: "Empty tree",
			tree: []interface{}{},
			want: true,
		},
		{
			name: "Simple symmetric",
			tree: []interface{}{1, 2, 2},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.tree)
			if got := IsSymmetric(root); got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		tree []interface{}
		want int
	}{
		{
			name: "Example 1",
			tree: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want: 3,
		},
		{
			name: "Example 2",
			tree: []interface{}{1, nil, 2},
			want: 2,
		},
		{
			name: "Empty tree",
			tree: []interface{}{},
			want: 0,
		},
		{
			name: "Single node",
			tree: []interface{}{1},
			want: 1,
		},
		{
			name: "Deep left tree",
			tree: []interface{}{1, 2, nil, 3, nil, 4},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.tree)
			if got := MaxDepth(root); got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name string
		tree []interface{}
		want int
	}{
		{
			name: "Example 1",
			tree: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want: 2,
		},
		{
			name: "Example 2",
			tree: []interface{}{2, nil, 3, nil, 4, nil, 5, nil, 6},
			want: 5,
		},
		{
			name: "Empty tree",
			tree: []interface{}{},
			want: 0,
		},
		{
			name: "Single node",
			tree: []interface{}{1},
			want: 1,
		},
		{
			name: "Balanced tree",
			tree: []interface{}{1, 2, 3},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.tree)
			if got := MinDepth(root); got != tt.want {
				t.Errorf("MinDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
