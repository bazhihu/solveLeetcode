package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

双指针+两个哑节点
分别将节点分别分布到两个哑节点上
然后合并两个哑节点

*/
func partition(head *ListNode, x int) *ListNode {
	var first = &ListNode{Val: 0}
	var second = &ListNode{Val: 0}
	firstHead := first
	secondHead := second

	for head != nil {
		tmp := head.Next
		if head.Val < x {
			first.Next = head
			first = first.Next
			first.Next = nil
		} else {
			second.Next = head
			second = second.Next
			second.Next = nil
		}
		head = tmp
	}

	if secondHead != nil {
		first.Next = secondHead.Next
	}
	return firstHead.Next
}
