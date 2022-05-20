package yee

import "strings"

type node struct {
	pattern  string  // 待匹配路由，只有叶子节点才有，其上为空，可用于判断是否匹配成功
	part     string  // 路由中的一部分，例如:lang，tries树的节点值
	children []*node // 子节点
	isWild   bool    // 是否精准匹配，当part有:或*时为true
}

// 匹配一层里第一个符合条件的叶子节点，用于插入
func (n *node) matchFirstChild(part string) *node {
	// 匹配节点的每个叶子节点
	for _, child := range n.children {
		if (part == child.part) || (child.isWild) {
			return child
		}
	}

	return nil
}

// 匹配一层里所用成功的节点，用于查找
func (n *node) machAllChild(part string) []*node {
	nodes := make([]*node, 0)

	for _, child := range n.children {
		if (part == child.part) || (child.isWild) {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	// 路径分片的长度等于高度
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchFirstChild(part)

	if child == nil {
		//当前匹配为空，创建新节点插入
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}

		return n
	}

	part := parts[height]
	children := n.machAllChild(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
