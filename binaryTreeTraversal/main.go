package main

import "fmt"

func main() {
	result := inorderTraversal(&TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	})
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	ret = inorder(root, ret)
	return ret
}

func inorder(node *TreeNode, a []int) []int {
	if node == nil {
		return a
	}
	a = inorder(node.Left, a)
	a = append(a, node.Val)
	a = inorder(node.Right, a)

	return a
}

func preorderTravesal(root *TreeNode) []int {
	ret := make([]int, 0)
	ret = preorder(root, ret)
	return ret
}

func preorder(node *TreeNode, a []int) []int {
	if node == nil {
		return a
	}
	a = append(a, node.Val)
	a = preorder(node.Left, a)
	a = preorder(node.Right, a)

	return a
}

func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	ret = postorder(root, ret)
	return ret
}

func postorder(node *TreeNode, a []int) []int {
	if node == nil {
		return a
	}
	a = postorder(node.Left, a)
	a = postorder(node.Right, a)
	a = append(a, node.Val)

	return a
}
