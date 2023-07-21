package addtwonumber

type ListNode struct {
      Val int
      Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    curr := result
    carry:= 0
    for l1!=nil||l2!=nil||carry>0{
        val:=0
        if l1!=nil{
            val += l1.Val
            l1=l1.Next
        }
        if l2!=nil{
            val += l2.Val
            l2=l2.Next
        }
        val += carry
        carry = val /10
        val = val%10
        curr.Val = val
        if l1!=nil || l2!=nil||carry>0{
            curr.Next = &ListNode{}
            curr = curr.Next
        }
    }
    return result
}