package gos

import (
	"strings"
)

type nodeType uint8

const (
	static nodeType = iota
	ROOT
	PARAM
	CATCHAll
)

type node struct {
	path        string
	nodeType    nodeType
	childrens   []*node
	handlerFunc HandlerFunc
	pathLen     int
	fullPath    string
}

// InitNewNode get a node
func initNewNode(path string, last bool, handlerFunc HandlerFunc) *node {
	children := new(node)
	if strings.Compare(string(path[0]), ":") == 0 {
		children.path = path[1:]
		children.nodeType = PARAM
	} else {
		children.path = path
		children.nodeType = CATCHAll
	}
	if last {
		children.handlerFunc = handlerFunc
	}
	return children
}

// GetPaths split url to path
func GetPaths(url string) []string {
	if strings.HasPrefix(url, "/") {
		url = string(url[1:])
	}
	if strings.HasSuffix(url, "/") {
		url = string(url[:len(url)-1])
	}
	if len(url) > 0 {
		return strings.Split(url, "/")
	} else {
		return make([]string, 0)
	}

}
func (n *node) AddRoute(url string, handlerFunc HandlerFunc) {
	paths := GetPaths(url)
	pathsLen := len(paths)
	fatherNode := n
	if pathsLen == 0 {
		if n.handlerFunc == nil {
			n.handlerFunc = handlerFunc
		} else {
			Error("duplicate route ", url)
		}
		return
	}
	for i, path := range paths {
		exist := false
		for _, cnode := range n.childrens {
			if strings.Compare(cnode.path, path) == 0 {
				exist = true
				fatherNode = cnode
				break
			}
			if cnode.nodeType == PARAM && strings.Compare(":"+cnode.path, path) == 0 {
				exist = true
				fatherNode = cnode
				break
			}
		}
		if i == pathsLen-1 {
			if exist {
				Error("duplicate route ", url)
			} else {
				cnode := initNewNode(path, true, handlerFunc)
				cnode.pathLen = pathsLen
				cnode.fullPath = url
				fatherNode.childrens = append(fatherNode.childrens, cnode)
			}
		}
		if !exist {
			cnode := initNewNode(path, false, handlerFunc)
			fatherNode.childrens = append(fatherNode.childrens, cnode)
			fatherNode = cnode
		}
	}
}

// SearchNodes search list match paths
func searchNodes(n *node, paths []string) []*node {
	nodes := make([]*node, 0, 10)
	if len(paths) == 0 {
		nodes = append(nodes, n)
		return nodes
	} else if len(paths) == 1 {
		for _, cnode := range n.childrens {
			if strings.Compare(cnode.path, paths[0]) == 0 {
				if cnode.handlerFunc != nil {
					nodes = append(nodes, cnode)
				}
				return nodes
			}
			if cnode.nodeType == PARAM {
				if cnode.handlerFunc != nil {
					nodes = append(nodes, cnode)
				}
			}
		}

	} else {
		for _, cnode := range n.childrens {
			if cnode.nodeType == PARAM || strings.Compare(paths[0], cnode.path) == 0 {
				nodes = append(nodes, searchNodes(cnode, paths[1:])...)
			}
		}
	}

	return nodes
}

// getMathcOne get one node
func getMatchOne(n *node, paths []string) *node {
	nodes := searchNodes(n, paths)
	if len(nodes) == 0 {
		return nil
	}

	for _, n := range nodes {
		if n.nodeType == CATCHAll {
			return n
		}
	}
	return nodes[0]

}
