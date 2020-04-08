package gee

type node struct {
	patter   string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node //子节点 例如：[doc, tutorial, intro]
	isWild   bool    //是否为精确匹配 part 含有 : 或 * 时为true
}
