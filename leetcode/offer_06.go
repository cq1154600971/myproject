package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	//ans := [n]int{}
	ans := make([]int, n)
	cur = head
	for cur != nil {
		ans[n - 1] = cur.Val
		cur = cur.Next
		n--
	}
	return ans
}
