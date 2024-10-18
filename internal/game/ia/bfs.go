package ia

import (
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

func BFS(curr, goal *point.Point) Stack[point.Point] {
	visited := []*point.Point{}

	root := &Node{
		point: curr,
	}
	queue := NewQueue[Node]()

	dirs := getDirections(curr)
	for _, p := range dirs {
		if world.Reachable(p) {
			node := &Node{
				point:  p,
				parent: root,
			}
			queue.Enqueue(node)
		}
	}

	for !queue.Empty() {
		node := queue.Dequeue()
		stop := false
		for _, v := range visited {
			if v.Equals(node.point) {
				stop = true
				break
			}
		}
		if stop {
			continue
		}
		if goal.Equals(node.point) {
			return node.traceBack(NewStack[point.Point]())
		}
		visited = append(visited, node.point)
		dirs := getDirections(node.point)
		for _, p := range dirs {
			if world.Reachable(p) {
				queue.Enqueue(&Node{
					point:  p,
					parent: node,
				})
			}
		}
	}

	return nil
}
