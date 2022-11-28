// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: 3

// Example 2:

// Input: root = [1,null,2]
// Output: 2

// Constraints:

//     The number of nodes in the tree is in the range [0, 104].
//     -100 <= Node.val <= 100

package leetcode

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	i := 0
	j := 0
	if root == nil {
		return 0
	}

	i = maxDepth(root.Left) + 1

	j = maxDepth(root.Right) + 1

	if i < j {
		return j
	} else if i > j {
		return i
	}

	return i
}
