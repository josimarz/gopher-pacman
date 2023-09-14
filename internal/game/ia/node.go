package ia

import (
	"github.com/josimarz/gopher-pacman/internal/game/point"
)

type Node struct {
	point  *point.Point
	parent *Node
}

func (n *Node) traceBack(path Stack[point.Point]) Stack[point.Point] {
	path.Push(n.point)
	if n.parent != nil {
		return n.parent.traceBack(path)
	}
	return path
}
