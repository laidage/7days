package gee

// 用来搜寻的，不是储存handler的结构
type Node struct {
	Pattern  string //用来映射handlers
	Children []*Node
	isWild   bool
}

func (node *Node) search() {

}

func (node *Node) insert() {

}
