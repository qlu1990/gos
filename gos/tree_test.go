package gos

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testRootNode = &node{
		nodeType: ROOT,
		path:     "",
		handlerFunc: func(c *Context) {
			fmt.Println("test middlerware ")
		},
		pathLen:   0,
		childrens: make([]*node, 0, 10),
	}
)

func TestInitNewNode(t *testing.T) {
	handlerFunc := func(c *Context) {
		fmt.Println("test middlerware ")
	}
	newNode := initNewNode("test", false, handlerFunc)
	if newNode.handlerFunc != nil || newNode.nodeType != CATCHAll {
		t.Error("error init new Node")
	}
	newNode = initNewNode(":test", false, handlerFunc)
	if newNode.handlerFunc != nil || newNode.nodeType != PARAM {
		t.Error("error init new Node")
	}
	newNode = initNewNode(":test", true, handlerFunc)
	if newNode.handlerFunc == nil || newNode.nodeType != PARAM {
		t.Error("error init new Node")
	}
	newNode = initNewNode("test", true, handlerFunc)
	if newNode.handlerFunc == nil || newNode.nodeType != CATCHAll {
		t.Error("error init new Node")
	}
}

func TestGetPaths(t *testing.T) {
	paths := GetPaths("/adc/aaa/ddd/")
	if strings.Compare(paths[0], "/") == 0 ||
		strings.Compare(paths[len(paths)-1], "/") == 0 {
		t.Error("error getPaths")
	}
}

func TestAddRoute(t *testing.T) {
	testRootNode.AddRoute("/url/test/:id", func(c *Context) {
		fmt.Println("/url/test/:id")
	})
	testRootNode.AddRoute("/url1/test", func(c *Context) {
		fmt.Println("/url1/test")
	})
	testRootNode.AddRoute("/:id/test", func(c *Context) {
		fmt.Println("/:id/test")
	})
	TraversingNode(testRootNode)
}

func TraversingNode(n *node) {
	for _, v := range n.childrens {
		if len(v.childrens) > 0 {
			TraversingNode(v)
		}
		if v.fullPath != "" {
			fmt.Println("path:" + v.fullPath)
		}

	}
}
func TestSearchNodes(t *testing.T) {
	Debug("start testSearchNodes")
	nodes := searchNodes(testRootNode, GetPaths("/url/test/aaa"))
	if len(nodes) == 0 {
		t.Error("error get nodes")
	}
	for _, v := range nodes {
		fmt.Println("search match: ", v.fullPath)
	}
	Debug("end testSearchNodes")
}
