package core

import (
	"fmt"
	"strings"
)

type (
	kind uint8
	// node is tree node
	node struct {
		kind     kind   // kind of node: common, param, wild
		fragment string // HTTP URL fragment
		handlers HandlersChain

		parent   *node
		children []*node
	}
	// tree is an alias for router
	tree struct {
		method string
		root   *node
	}
	// MethodForest is alias for MethodTrees
	MethodForest []*tree
)

const (
	// root kind
	root kind = iota
	// common kind
	cKind
	// param kind
	pKind
	// wild kind
	wKind
	nullString = ""
	charSlash  = '/'
	strSlash   = "/"
	charColon  = ':'
	charStar   = '*'
	strColon   = ":"
	strStar    = "*"
)

// get method tree according to the HTTP method, ok will be true if exists
func (forest MethodForest) get(method string) (router *tree, ok bool) {
	for _, tree := range forest {
		if tree.method == method {
			return tree, true
		}
	}
	return nil, false
}

// addRoute adds a node with the given handle to the path
func (t *tree) addRoute(path string, handlers HandlersChain) {
	validatePath(path)
	// purePath is path without any handle
	purePath := path
	if handlers == nil {
		panic(fmt.Sprintf("Adding route without handler function: %v", purePath))
	}
	t.insert(purePath, handlers)
}

// insert into trie tree
func (t *tree) insert(path string, handlers HandlersChain) {
	if t.root == nil {
		t.root = &node{
			kind:     root,
			fragment: strSlash,
		}
	}
	currNode := t.root
	fragments := strings.Split(path, strSlash)[1:]
	for i, fragment := range fragments {
		child := currNode.matchChild(fragment)
		if child == nil {
			child = &node{
				kind:     matchKind(fragment),
				fragment: fragment,
				parent:   currNode,
			}
			currNode.children = append(currNode.children, child)
		}
		if i == len(fragments)-1 {
			child.handlers = handlers
		}
		currNode = child
	}
}

func (t *tree) matchBranch() {

}

// matchChild for the node children field with HTTP URL fragment
func (n *node) matchChild(fragment string) *node {
	for _, child := range n.children {
		if child.fragment == fragment {
			return child
		}
	}
	return nil
}

// matchKind according to the HTTP URL fragment
func matchKind(fragment string) kind {
	switch fragment[0] {
	case charColon:
		return pKind
	case charStar:
		return wKind
	default:
		return cKind
	}
}
