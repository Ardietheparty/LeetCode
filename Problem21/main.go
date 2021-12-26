package main

import (
	"fmt"
	"time"
)
/*
21. Merge Two Sorted Lists
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func main() {
	timeStart := time.Now()
	//a := []int{1,2,3,4}
	//b := []int{1,2,2,6}
	//x := makeList(a)
	//y := makeList(b)
	//z := mergeTwoLists(x,y)
	//q := meTwoLists(x,y)
	//z.PrintNode()
	fmt.Println(" ")
	fmt.Println(" ")
	//q.PrintNode()
	c1 := ListNode{1,nil}
	maker := ListNode{2,nil}
	n  := c1
	n.Next = &maker
	c1 = n
	mak := ListNode{3,nil}
	n = c1
	c1= c1.Next

	c1=n
	c1.PrintNode()



	fmt.Println(time.Since(timeStart))
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	a:=0
	if list1.Val < list2.Val {
		a = list1.Val
		list1 = list1.Next
	} else {
		a = list2.Val
		list2 = list2.Next
	}
	out := ListNode{a,nil}
	for list1 != nil || list2 != nil {
		if list1 == nil {
			a = list2.Val
			//out.AddNode(list2.Val)
			list2=list2.Next
		} else if list2 == nil {
			a = list1.Val
			//out.AddNode(list1.Val)
			list1 = list1.Next
		} else if list1.Val<list2.Val{
			a = list1.Val
			//out.AddNode(list1.Val)
			list1 = list1.Next
		} else {
			a = list2.Val
			//out.AddNode(list2.Val)
			list2=list2.Next
		}
		newNode := ListNode{a, nil}
		iter := out
		for iter.Next != nil {
			iter = iter.Next
		}
		iter.Next = &newNode
	}
	return &out
}
func meTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	a := add(list1)
	b := add(list2)
	c := merg(a,b)

	return makeList(c)
}
func add(l1 *ListNode) []int {
	c1 := l1
	var out []int
	for c1 != nil {
		out = append(out, c1.Val)
		c1 = c1.Next
	}

	return out
}
func (n *ListNode) AddNode(Val int) {
	newNode := ListNode{Val, nil}
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

func makeList(a []int) *ListNode {
	list := ListNode{a[0],nil}
	for i := 1; i < len(a); i++ {
		list.AddNode(a[i])
	}
	return &list
}
func merg(a []int, b []int) []int {
	A := len(a)
	B := len(b)
	C := A+B
	j,k := 0,0
	var out []int
	for i := 0; i < C; i++ {
		if j >= A {
			out = append(out, b[k])
			k++
		} else if k >= B {
			out = append(out,a[j])
			j++
		} else {
			m := a[j]
			n := b[k]
			if m < n {
				out = append(out,m)
				j++
			} else {
				out = append(out,n)
				k++
			}
		}
	}
	return out
}