/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 var sum int

 func sumOfLeftLeaves(root *TreeNode) int {
	 sum = 0
	 if root == nil {
		 return 0
	 }
	 if root.Right == nil && root.Left == nil {
		 return 0
	 }
	 
	 addLeftTree(root, 1) 
	 
	 return sum
 }
 
 func addLeftTree(node *TreeNode, flag int) {
	 // check if node exists
	 if node == nil {
		 return
	 }
	 // go left if we can
	 addLeftTree(node.Left, 1)
	 
	 // add val if it comes from left leaf and no children
	 if node.Right == nil && node.Left == nil && flag == 1 {
		 sum += node.Val
	 }   
	 
	 // check right side
	 addLeftTree(node.Right, 0)
 
 }