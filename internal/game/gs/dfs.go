package gs

import (
	"github.com/josimarz/gopher-pacman/internal/game/stack"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

type DepthFirstSearch struct {
	explored []*tile.Point
}

func NewDepthFirstSearch() *DepthFirstSearch {
	return &DepthFirstSearch{}
}

func (s *DepthFirstSearch) Run(pos, goal *tile.Point) *stack.Stack[tile.Point] {
	root := &Node{
		point: pos,
	}
	st := stack.New[Node]()
	st.Push(root)

	for !st.Empty() {
		node := st.Pop()

		if node.point.Equals(goal) {
			return node.traceBack(stack.New[tile.Point]())
		}

		neighbors := []*tile.Point{
			node.point.Up(),
			node.point.Down(),
			node.point.Left(),
			node.point.Right(),
		}

		for _, p := range neighbors {
			if t := world.Instance().TileAt(p); s.reachable(t) {
				s.visit(t.Point())
				st.Push(&Node{
					point:  p,
					parent: node,
				})
			}
		}
	}

	return nil
}

func (s *DepthFirstSearch) visit(p *tile.Point) {
	s.explored = append(s.explored, p)
}

func (s *DepthFirstSearch) reachable(t *tile.Tile) bool {
	return t != nil && !s.visited(t.Point()) && t.Accessible()
}

func (s *DepthFirstSearch) visited(p *tile.Point) bool {
	for _, v := range s.explored {
		if v.Equals(p) {
			return true
		}
	}
	return false
}
