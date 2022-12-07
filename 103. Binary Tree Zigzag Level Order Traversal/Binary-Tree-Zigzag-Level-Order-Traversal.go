package leetcode

// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
// (i.e., from left to right, then right to left for the next level and alternate between).

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[20,9],[15,7]]

// Example 2:

// Input: root = [1]
// Output: [[1]]

// Example 3:

// Input: root = []
// Output: []

// Constraints:

//     The number of nodes in the tree is in the range [0, 2000].
//     -100 <= Node.val <= 100

//   Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
 }
 

type NodeWithLevel struct {
    node *TreeNode
    level int
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil{
        return [][]int{}
    }
    q := []NodeWithLevel{
        {
        node: root,
        level: 0,
        },
    }
    visited := []*TreeNode{}
    res := [][]int{{root.Val}}
    tmp := []int{}
    lev := 0

    for len(q) > 0{
        vertex := q[0]
        node, level := vertex.node, vertex.level
        visited = append(visited, node)
        
        q = q[1:]
        if lev < level{
            lev = level
            if lev % 2 != 0 && len(tmp) > 1{
                tmp = revers(tmp)
                res = append(res, tmp)
            } else {              
            res = append(res, tmp)
            }
            tmp = []int{}
        }
        
        if node.Left != nil{
            leftNode := NodeWithLevel{
                node: node.Left,
                level: level+1,
            }
            q = append(q, leftNode)
            tmp = append(tmp, leftNode.node.Val)
        }
        if node.Right != nil{
            rightNode := NodeWithLevel{
                node: node.Right,
                level: level+1,
            }
            q = append(q, rightNode)
            tmp = append(tmp, rightNode.node.Val)
        }   
        
        }
    
    return res
}

func revers(tmp []int) []int{
   var res []int

   for i := len(tmp) - 1; i >= 0; i-- {
      res = append(res, tmp[i])
   }

   return res
}