package gee

import (
	"strings"
	"log"
)

// test: /hello/one /hello/two /:lang/ab /:lang/cd
// 用来搜寻的，不是储存handler的结构
type Node struct {
	Pattern  string //用来映射handlers
	Children []*Node
	isWild   bool
	Path     string
}

func NewNode() *Node {
	return &Node{
		Pattern:  "",
		Children: make([]*Node, 0),
	}
}

func (node *Node) search(parts []string, height int) *Node {
	if height >= len(parts) {
		return nil
	}
	for _, child := range node.Children {
		log.Printf("%v", parts[height])
		if parts[height] == child.Path || child.isWild == true {
			log.Printf("%v", parts[height])
			var resultNode *Node
			if height == len(parts) - 1 {
				log.Printf("%v22", parts[height])
				return child
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
				child.Pattern = strings.Join(parts, "/")
				log.Printf("%v", child.Pattern)
			} else {
				child.insert(parts, height+1)
			}
			break
		}
	}
	if flag == -1 {
		childNode := NewNode()
		childNode.Path = parts[height]
		if strings.Contains(parts[height], "*") || strings.Contains(parts[height], ":") {
			childNode.isWild = true
		} else {
			childNode.isWild = false
		}
		if height == len(parts)-1 {
			childNode.Pattern = strings.Join(parts, "/")
			log.Printf("%v", childNode.Pattern)
		} else {
			childNode.insert(parts, height+1)
		}
		node.Children = append(node.Children, childNode)
	}
}
