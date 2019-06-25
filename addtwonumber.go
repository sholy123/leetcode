package main


type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0,nil}
	sh := head
	carry := 0
	current := 0
	for l1 != nil || l2 != nil || carry > 0{
		sum := current
		if l1 != nil{
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current = (sum + carry) % 10
		sh.Next = new(ListNode)
		sh.Next.Val = current
		sh = sh.Next
	}
	return head.Next
}
func main(){

}