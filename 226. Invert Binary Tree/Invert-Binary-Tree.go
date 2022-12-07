package leetcode

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    q := []*TreeNode{root}

    for len(q) > 0{
        node := q[0]
        q = q[1:]

        node.Left, node.Right = node.Right, node.Left

        if node.Left != nil {
            q = append(q, node.Left)
        }

        if node.Right != nil {
            q = append(q, node.Right)
        }

    }

    return root
    
}
