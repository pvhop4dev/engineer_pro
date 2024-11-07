package leetcode

func KthSmallest(root *TreeNode, k int) int {
	inOrder(root, k)
	return arr[k-1]
}

var arr []int

func inOrder(root *TreeNode, k int) {
	if root == nil {
		return
	}
	inOrder(root.Left, k)
	arr = append(arr, root.Val)
	inOrder(root.Right, k)
}
