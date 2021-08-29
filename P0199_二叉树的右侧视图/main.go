package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		res = append(res, queue[levelSize-1].Val)
		for i := 0; i < levelSize; i ++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

func main(){
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode {
		Val: 2,
	}
	node3 := &TreeNode {
		Val: 3,
	}
	node4 := &TreeNode {
		Val: 4,
	}
	node5 := &TreeNode {
		Val: 5,
	}
	node6 := &TreeNode {
		Val: 6,
	}
	node7 := &TreeNode {
		Val: 7,
	}
	root := &TreeNode{
		Val: 0,
	}
	root.Left = node1
	root.Right = node2
	root.Left.Left = node3
	root.Right.Right = node4
	root.Right.Right.Right = node5
	root.Right.Right.Left = node6
	root.Right.Right.Left.Left = node7
	res := rightSideView(root)
	res1 := rightSideView(nil)
	fmt.Printf("res=%v, res1=%v\n", res, res1)
}