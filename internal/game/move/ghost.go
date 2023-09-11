package move

import (
	"math/rand"
	"time"

	"github.com/josimarz/gopher-pacman/internal/game/direction"
	"github.com/josimarz/gopher-pacman/internal/game/gs"
	"github.com/josimarz/gopher-pacman/internal/game/stack"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

var (
	r *rand.Rand
)

type GhostTracking struct {
	path      *stack.Stack[tile.Point]
	dir       direction.Direction
	currPoint *tile.Point
	nextPoint *tile.Point
	speed     int
}

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	r = rand.New(src)
}

func NewGhostTracking(pt *tile.Point) *GhostTracking {
	return &GhostTracking{
		path:      stack.New[tile.Point](),
		dir:       direction.Up,
		currPoint: pt.Clone(),
		nextPoint: pt.Clone(),
		speed:     1,
	}
}

func (t *GhostTracking) CurrPoint() *tile.Point {
	return t.currPoint
}

func (t *GhostTracking) Dir() direction.Direction {
	return t.dir
}

func (t *GhostTracking) Move() {
	if t.currPoint.Equals(t.nextPoint) {
		if t.path == nil || t.path.Empty() {
			t.recreatePath()
		}
		t.nextPoint = t.path.Pop()
		t.dir = t.currPoint.Dir(t.nextPoint)
	}
	t.moveX()
	t.moveY()
}

func (t *GhostTracking) recreatePath() {
	var goal *tile.Point
	var content tile.Content
	for {
		x := r.Intn(21)
		y := r.Intn(21)
		content = world.ContentSet[y][x]
		goal = tile.NewPoint(x*tile.Size, y*tile.Size)
		if goal.Equals(t.currPoint) {
			continue
		}
		if content != tile.Outline && content != tile.Wall {
			break
		}
	}
	dfs := gs.NewDepthFirstSearch()
	t.path = dfs.Run(t.currPoint.Clone(), goal)
}

func (t *GhostTracking) moveX() {
	if t.currPoint.X > t.nextPoint.X {
		t.currPoint.X -= t.speed
	} else if t.currPoint.X < t.nextPoint.X {
		t.currPoint.X += t.speed
	}
}

func (t *GhostTracking) moveY() {
	if t.currPoint.Y > t.nextPoint.Y {
		t.currPoint.Y -= t.speed
	} else if t.currPoint.Y < t.nextPoint.Y {
		t.currPoint.Y += t.speed
	}
}
