package main

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

示例:

输入: 1->2->NULL, m = 1, n = 2
输出: 2->1->NULL

 * 1 增加哨兵 防止是完全逆转
 * 2 先找到要逆转的前一个节点
 * 3 然后把要逆转的下个节点一直搬运到伪哨兵的Next节点位置
 * 4 完成区间逆转

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var first = &ListNode{Val: 0}
	first.Next = head
	tmp := first
	for i := 1; i < m; i++ { // 找到M的前一个节点 （M=1的话增加哨兵）
		tmp = tmp.Next
	}

	head = tmp.Next //该节点一直移动到链表 区间的最后 无需递进
	for i := m; i < n; i++ {
		t := head.Next     //取该节点的Next节点 将要换到区间首部
		head.Next = t.Next // 当前节点取代Next节点位置

		t.Next = tmp.Next // 将t节点 插入到tmp 中间去
		tmp.Next = t
	}

	return first.Next
}
