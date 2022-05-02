package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  7,
				Left: nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	fmt.Println(kthSmallest(tree, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	sorted := inorder(root, []int{})
	return sorted[k-1]
}

// TODO: no need to traverse to all node, try to break traversing and return the kth smallest value instantly
func inorder(tree *TreeNode, sorted []int) []int {
	if tree == nil {
		return sorted
	}

	sorted = inorder(tree.Left, sorted)
	sorted = append(sorted, tree.Val)
	sorted = inorder(tree.Right, sorted)

	return sorted
}
