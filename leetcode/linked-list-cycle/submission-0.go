// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle
//
// Using turtoise and hare, still dont know why this work btw: https://cp-algorithms.com/others/tortoise_and_hare.html
package linkedlistcycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}
