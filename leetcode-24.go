package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val int
    Next *ListNode
}
func Swap(p,q *ListNode){
	temp := new(ListNode)
	temp = q
	q = p
	p = temp
}
func swapPairs(head *ListNode) *ListNode {
	if (head == nil || head.Next ==nil){
		return head
	}
	L := head
	head = head.Next
	var q *ListNode
	var p *ListNode
	for head != nil && head.Next != nil{
		q = head.Next
		p = head
		Swap(p,q)
		head = q.Next
	}
	return L
}
func Add(head *ListNode,data int){
	q := head
	for q!=nil{
		q = q.Next
	}
	q.Val = data
}
func main(){
	head := new(ListNode)
	q := head
	for i:=0; i < 5;i++{
		Add(q,i)
	}
	for head!=nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
