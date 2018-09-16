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
	num := getNodeNum(root)
	fmt.Printf("node num = %d\n",num)

}