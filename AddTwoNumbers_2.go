package LeetCode_Go

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret, adder *ListNode

	carry := 0
	for ;; {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		num1 := 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		num2 := 0
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		sum := num1 + num2 + carry
		carry = sum / 10
		sum = sum % 10

		if adder == nil {
			adder = &ListNode{sum, nil}
			ret = adder
		} else {
			adder.Next = &ListNode{sum, nil}
			adder = adder.Next
		}
	}

	return ret
}