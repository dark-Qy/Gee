package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由
	part     string  // 路由中的一部分
	children []*node // 子节点
	isWild   bool    // 是否精确匹配，part含有:或*时为true
}

// 理顺一下思路：首先对于路由来说，最重要的就是两个功能：注册和查找。
// 注册的过程就是将路由添加到Trie树中，查找的过程就是在Trie树中找到与之匹配的路由。

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 注册路由
func (n *node) insert(pattern string, parts []string, height int) {
	// 如果到底了
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 查找子节点是否有跟part匹配的
	part := parts[height]
	child := n.matchChild(part)

	// 如果没有，新建一个
	if child == nil {
		child = &node{part: part, isWild: part[0] == '*' || part[0] == ':'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 查找路由
func (n *node) search(parts []string, height int) *node {
	// 如果到底了
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		// 递归查找
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
