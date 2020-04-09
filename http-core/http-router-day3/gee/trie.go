package gee

type node struct {
	patter   string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node //子节点 例如：[doc, tutorial, intro]
	isWild   bool    //是否为精确匹配 part 含有 : 或 * 时为true
}

//第一个匹配成功的节点 用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//匹配所有成功的节点 用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
}
