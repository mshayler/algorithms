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