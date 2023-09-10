package gs

import (
	"github.com/josimarz/gopher-pacman/internal/game/stack"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
)

type Node struct {
	point  *tile.Point
	parent *Node
}

func (n *Node) traceBack(path *stack.Stack[tile.Point]) *stack.Stack[tile.Point] {
	path.Push(n.point)
	if n.parent != nil {
		return n.parent.traceBack(path)
	}
	return path
}
