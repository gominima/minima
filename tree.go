package minima

import (
	"regexp"
	"strings"
)

// tree is a trie tree.
type tree struct {
	node *node
}

// node is a node of tree.
type node struct {
	path     string
	actions  map[string]*action // key is method
	children map[string]*node   // key is label of next nodes
}

// action is an action.
type action struct {
	handlers []Handler
	Param    []string
	regex    *regexp.Regexp
}

type result struct {
	action *action
}

func NewTree() *tree {
	return &tree{
		node: &node{
			path:     "/",
			actions:  make(map[string]*action),
			children: make(map[string]*node),
		},
	}
}

func (t *tree) InsertNode(method string, path string, handlers ...Handler) error {
	reg, Params := RegexPath(path)
	n := t.node
	if path == "/" {
		n.path = path
		n.actions[method] = &action{
			handlers: handlers,
			Param:    Params,
			regex:    regexp.MustCompile(reg),
		}
		return nil
	}
	p := t.parsepath(path)
	for i, v := range p {
		nextnode, ok := n.children[v]
		if ok {
			n = nextnode
		}
	}
}

func (t *tree) parsepath(str string) []string {
	s := strings.Split(str, "/")
	var rval []string
	for _, str := range s {
		if str != "" {
			rval = append(rval, str)
		}
	}
	return rval
}
