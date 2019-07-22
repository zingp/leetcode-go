package main

import (
	"fmt"
)

// ListNode 链表结构体定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 判断是否是回文链表
func isPalindrome(head *ListNode) bool {
	// 如果只有一个元素
	if head == nil || head.Next == nil {
		return true
	}

	// 用两个指针，慢指针每次移动一位，快指针每次移动两位，找到中点slow
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转后半段链表，pre就是反转后的头结点
	var pre *ListNode
	var next *ListNode
	for slow != nil {
		next = slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}

	// 对比反转后的链表和未反转的链表中的每一个元素
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}

	return true
}

// 单链表反转
// 用两个指针，pre代表当前节点的前一个节点，next代表当前节点的后一个节点
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// 打印链表的每一个元素
func rangeList(head *ListNode) {
	if head == nil {
		fmt.Println(nil)
	}
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// 寻找链表中点的节点
func selectMidListNode(head *ListNode) *ListNode {
	// 如果只有一个元素
	if head == nil {
		return head
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 寻找链表中点，且偶数时返回中间2个元素的第一个元素
func selectMidListNode2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	if fast.Next != nil || fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 在O(1)时间复杂度内删除单链表指定节点
func deleteNode(head *ListNode, node *ListNode) *ListNode {
	// 如果输入参数有空，直接返回头结点
	if head == nil || node == nil {
		return head
	}
	// 如果删除的是头结点
	if head == node {
		return head.Next
	}

	//如果链表有两个及以上节点，且删除的是尾节点
	if node.Next == nil {
		// 找待删除元素的前驱节点
		tmp := head
		for tmp.Next != node {
			tmp = tmp.Next
		}
		// 删除节点
		tmp.Next = nil
	} else {
		// 如果链表有两个及以上节点，且删除的是某个中间结点
		// 将待删除的节点的下一个节点的值 付给待删除的节点
		node.Val = node.Next.Val
		// 将待删除节点的下节点指向下下个节点，即删除待删除节点的下一个节点
		node.Next = node.Next.Next
	}
	// 返回删除节点后的链表头结点
	return head
}

// 反转部分单链表
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	i := 1
	var reverseNewHead *ListNode  // 反转部分链表反转后的头结点，反转过程中这个节点不断变化
	var reverseNewTail *ListNode  // 第m个节点，即反转部分链表反转后的尾结点
	var reversePreNode *ListNode  // 第m-1 个节点，即反转部分链表反转前其头结点的前一个结点
	var reverseNextNode *ListNode // 反转时，用于挂住下一个节点的节点，当前head的下一个节点
	oldHead := head
	for head != nil {
		if i > n {
			break
		}
		if i == m-1 {
			reversePreNode = head
		}
		if i >= m && i <= n {
			if i == m {
				reverseNewTail = head
			}
			reverseNextNode = head.Next
			head.Next = reverseNewHead
			reverseNewHead = head
			head = reverseNextNode
		} else {
			head = head.Next
		}
		i++
	}
	// 反转部分反转之后，整体反转一次，连接断开的地方
	// 反转部分反转之后，第m个节点的下一个节点应该是第n+1个节点
	reverseNewTail.Next = reverseNextNode
	// 第m-1个节点不是空节点，则他的下一个节点应该反转部分的头结点
	if reversePreNode != nil {
		reversePreNode.Next = reverseNewHead
		return oldHead
	}
	// 如果第m-1个节点是空节点说明m=1，从头开始反转；新的头结点就是reverseNewHead
	return reverseNewHead
}

// 旋转单链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 1
	cur := head
	// 注意如果这里是cur != nil 最后cur就是nil了
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	// 首尾连接
	cur.Next = head
	m := n - k%n
	for i := 0; i < m; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	return newHead
}

//  删除单链表倒数第 k 个节点
func deleteKNode(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	i := 0
	for fast != nil {
		fast = fast.Next
		if i >= k {
			slow = slow.Next
		}
		i++
	}
	slow.Val = slow.Next.Val
	slow.Next = slow.Next.Next

	return head
}

// 单链表划分
// 样例：给定链表 1->4->3->2->5->2->null，并且 x=3，返回 1->2->2->4->3->5->null
// 新建两个链表left,right;left存放小于x的，right存放大于等于x的
// 最后将left尾节点与right头结点连接起来
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var leftDummy = &ListNode{}
	var rightDummy = &ListNode{}
	left := leftDummy
	right := rightDummy

	for head != nil {
		if head.Val < x {
			left.Next = head
			left = head
		} else {
			right.Next = head
			right = head
		}
		head = head.Next
	}
	right.Next = nil
	left.Next = rightDummy.Next
	return leftDummy.Next
}

// 链表求和
// 样例：给出两个链表 3->1->5->null 和 5->9->2->null，返回 8->0->8->null
func addList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	pre := &ListNode{}
	cur := pre
	add1 := 0
	for l1 != nil {
		s := l1.Val + l2.Val + add1
		cur.Next = &ListNode{
			Val: s % 10,
		}
		cur = cur.Next
		add1 = s / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	if add1 == 1 {
		cur.Next = &ListNode{
			Val: 1,
		}
	}
	return pre.Next
}

// 快速排序
func quickSortList(head *ListNode, end *ListNode) {
	if head != end {
		node := partitionList(head, end)
		quickSortList(head, node)
		quickSortList(node.Next, end)
	}
}

// 快排中的一次分割
func partitionList(head *ListNode, end *ListNode) *ListNode {
	p1 := head
	p2 := head.Next

	for p2 != end {
		if p2.Val < head.Val {
			p1 = p1.Next
			fmt.Printf("p1.Val=%d, p2.Val=%d\n", p1.Val, p2.Val)
			// 此时，p1 指向的是小于head.Val的节点；p2指向的是大于head.Val的节点；故相互调换
			p1.Val, p2.Val = p2.Val, p1.Val
		}
		p2 = p2.Next
	}

	if p1 != head {
		p1.Val, head.Val = head.Val, p1.Val
	}
	return p1
}

// 测试快排
func testQuickSort() {

	n1 := &ListNode{Val: 8}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 7}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5, Next: nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	quickSortList(n1, nil)
	rangeList(n1)
}

// 合并两个有序链表
func mergeSortList(head1, head2 *ListNode) *ListNode {

	return head1
}

// 切片转链表
func silce2List(s []int) *ListNode {
	pre := &ListNode{}
	r := pre
	for _, v := range s {
		r.Next = &ListNode{
			Val: v,
		}
		r = r.Next
	}
	return pre.Next
}

// 生成测试链表
func genListNode() *ListNode {
	n1 := &ListNode{Val: 8}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 7}
	n4 := &ListNode{Val: 4}
	// n5 := &ListNode{Val: 5, Next: nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	// n4.Next = n5

	return n1
}

func main() {
	// head := genListNode()
	// rangeList(head)
	// fmt.Println("----------------------")
	// mid := selectMidListNode2(head)
	// fmt.Printf("mid:%d", mid.Val)

	// rangeList(deleteKNode(head, 2))
	// rangeList(rotateRight(head, 3))
	// rangeList(reverseBetween(head, 2, 4))
	// i := 1
	// cur := head
	// for cur != nil {
	// 	if i == 4{
	// 		break
	// 	}
	// 	cur = cur.Next
	// 	i++
	// }
	// // cur 就是第4个节点
	// rangeList(deleteNode(head, cur))

	// rangeList(partition(head, 5))
	// s1 := []int{6, 2, 3, 9}
	// s2 := []int{4, 3, 2, 1}
	// l1 := silce2List(s1)
	// l2 := silce2List(s2)
	// rangeList(addList(l1, l2))

	testQuickSort()
}
