package main

import (
	"fmt"
)

// TreeNode is binary tree Node
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 获取二叉树节点数
func getNodeNum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getNodeNum(root.Left) + getNodeNum(root.Right) + 1
}

// 获取二叉树的最大深度
func maxDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}

	x := maxDeep(root.Left)
	y := maxDeep(root.Right)
	if x > y {
		return x + 1
	} 
	return y + 1
}

// 获取二叉树最小深度
func minDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDeep(root.Left)
	right := minDeep(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	if left > right {
		return right + 1
	}
	return left + 1
}

// 先序遍历
func preOrderTraversal(root *TreeNode, r []int) []int {
	if root == nil {
		return r
	}
	r = append(r, root.Val)
	r = preOrderTraversal(root.Left, r)
	r = preOrderTraversal(root.Right, r)
	return r
}

func geneBinaryTree() *TreeNode {
	n1 := &TreeNode{Val:1,}
	n2 := &TreeNode{Val:2,}
	n3 := &TreeNode{Val:3,}
	n4 := &TreeNode{Val:4,}
	n5 := &TreeNode{Val:5,}
	n6 := &TreeNode{Val:6,}
	n7 := &TreeNode{Val:7,}
	n8 := &TreeNode{Val:8,}
	n9 := &TreeNode{Val:9,}
	n10 := &TreeNode{Val:10,}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	n4.Left = n8
	n4.Right = n9
	n5.Left = n10
	return n1
}

func main(){

	root := geneBinaryTree()
	// num := getNodeNum(root)
	// fmt.Printf("node num = %d\n",num)

	// fmt.Printf("max deep:%d", maxDeep(root))

	// fmt.Printf("min deep:%d", minDeep(root))
	var r []int
	fmt.Printf("pre order traversql:%v", preOrderTraversal(root, r))
}