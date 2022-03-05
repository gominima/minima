package minima

import (
	"sync"
)

const (
	Tsafe = 1 << iota
	// Tdebug adds more information to the tree's string representation.
	Tdebug
	// Tbinary uses a binary PATRICIA tree instead of a prefix tree.
	Tbinary
)

type tree struct {
	root        *Node
	len         int
	size        int
	safe        bool
	placeholder byte
	delim       byte
	mu          *sync.Mutex
}

func NewTree() *tree {
	return &tree{
		root: &Node{},
		len: 1,
		mu: &sync.Mutex{},
		safe: true,
	}
}

func (tr *tree) InsertNode(key string, handler Handler) {
	if key == "" || handler == nil {
		return
	}
	if tr.safe {
		defer tr.mu.Unlock()
		tr.mu.Lock()
	}
	
}