package easy

// https://leetcode.cn/problems/binary-tree-paths/

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var path []string
	var res []string
	traverse(root, &path, &res)
	return res
}

func traverse(root *TreeNode, path *[]string, res *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*path = append(*path, strconv.Itoa(root.Val))
		*res = append(*res, strings.Join(*path, "->"))
		*path = (*path)[:len(*path)-1]
		return
	}
	*path = append(*path, strconv.Itoa(root.Val))
	traverse(root.Left, path, res)
	traverse(root.Right, path, res)
	*path = (*path)[:len(*path)-1]
}
