/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
// Leetcode 2
// Rather than popping or storing, this approach uses basic math principles
// We create the start of a linked list with an empty array
// Then create a Next node for each loop, that adds the values from the other two if there is one, OR the value of the carryover from the last add
// return the list, with the one after.


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    
    sum := 0
    curr := head
    for {
        
        // add val to sum for each, and then move the pointer forward.
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        
        
        // Create a new node with the remainder value, then move our pointer to it
        curr.Next = &ListNode{
            Val: sum % 10,
        }
        curr = curr.Next
        
        
        
        // divide sum by 10 for each loop
        sum /= 10
        
        // check if either nodes next isnt nil, or if we have some remainder that needs to be included, that is in sum
        if l1 != nil || l2 != nil || sum > 0 {
            continue
        }
        break
    }
    
    // we started with a blank node, so return head.Next
    return head.Next
    
}