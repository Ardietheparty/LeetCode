package main

import (
	"fmt"
	"time"
)
/*
Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func main() {
	timeStart := time.Now()
	a := *ListNode{2,nil}
	b := *ListNode{5,nil}
	a.AddNode(4)
	a.AddNode(3)
	b.AddNode(6)
	b.AddNode(4)
	a.PrintNode()
	b.PrintNode()
	c := addTwoNumbers(a,b)
	c.PrintNode()
	fmt.Println(time.Since(timeStart))
}
func (n *ListNode) AddNode(data int) {
	newNode := ListNode{data,nil}
	iter := n
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &newNode
}
func (n *ListNode) PrintNode() {
	iter := n
	for iter != nil {
		fmt.Println(iter.Val)
		iter = iter.Next
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var c int
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || c != 0 {
		a,b := 0,0
		if l1 != nil {
			a,l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			b,l2 = l2.Val, l2.Next
		}
		n := a + b + c
		c = n/10
		cur.Next = &ListNode{n%10,nil}
		cur = cur.Next
	}
	return head.Next
}