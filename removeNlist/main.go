package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmpNode := head
	lenHead := 0
	for tmpNode != nil{
		lenHead++
		tmpNode = tmpNode.Next
	}
	difference := lenHead - n

	if difference != 0 {
		tmpNode = head
		for i := 0; i < difference - 1; i++{
			tmpNode = tmpNode.Next
		}
		tmpNode.Next = tmpNode.Next.Next
	}
	if lenHead == 1 && n == 1{
		head = nil
	}
	if lenHead == n{
		head = head.Next
	}
	return head
}
