package gee

import (
	"strings"
)

// test: /hello/one /hello/two /:lang/ab /:lang/cd
// 用来搜寻的，不是储存handler的结构
type Node struct {
	Pattern  string //用来映射handlers
	Children []*Node
	isWild   bool
	Path     string
}

func newNode() *Node {
	return &Node{
		Pattern:  "",
		Children: make([]*Node, 0),
		isWild:   false,
	}
}

func (node *Node) search(parts []string, height int) *Node {
	for _, child := range node.Children {
		if parts[height] == child.Path || child.isWild == true {
			var resultNode *Node
			if height == len(parts)-1 || strings.HasPrefix(child.Path, "*") {
				if child.Pattern == "" {
					return nil
				} else {
					return child
				}
			} else {
				resultNode = child.search(parts, height+1)
			}
			if resultNode != nil {
				return resultNode
			}
		}
	}
	return nil
}

// 插入["get"] ["post"]试试
func (node *Node) insert(parts []string, height int) {
	flag := -1
	for index, child := range node.Children {
		if child.Path == parts[height] {
			flag = index
			if height == len(parts)-1 {
				child.Pattern = strings.Join(parts, "-")
			} else {
				child.insert(parts, height+1)
			}
			break
		}
	}
	if flag == -1 {
		childNode := newNode()
		childNode.Path = parts[height]
		if strings.Contains(parts[height], "*") || strings.Contains(parts[height], ":") {
			childNode.isWild = true
		}
		if height == len(parts)-1 {
			childNode.Pattern = strings.Join(parts, "-")
		} else {
			childNode.insert(parts, height+1)
		}
		node.Children = append(node.Children, childNode)
	}
}
