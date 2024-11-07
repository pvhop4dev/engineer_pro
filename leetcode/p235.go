package leetcode

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		return LowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return LowestCommonAncestor(root.Right, p, q)
	}
	return root
}
