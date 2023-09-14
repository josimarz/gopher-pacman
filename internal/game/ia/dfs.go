package ia

import (
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

func DepthFirstSearch(curr, goal *point.Point) Stack[point.Point] {
	visited := []*point.Point{}
	root := &Node{
		point: curr,
	}
	st := NewStack[Node]()
	st.Push(root)

	for !st.Empty() {
		node := st.Pop()

		if node.point.Equals(goal) {
			return node.traceBack(NewStack[point.Point]())
		}

		directions := getDirections(node.point)

		for _, p := range directions {
			if world.Reachable(p) {
				found := false
				for _, v := range visited {
					if v.Equals(p) {
						found = true
						break
					}
				}
				if found {
					continue
				}
				visited = append(visited, p)
				st.Push(&Node{
					point:  p,
					parent: node,
				})
			}
		}
	}
	return nil
}
