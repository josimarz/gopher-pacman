package move

import (
	"math/rand"
	"time"

	"github.com/josimarz/gopher-pacman/internal/game/direction"
	"github.com/josimarz/gopher-pacman/internal/game/event"
	"github.com/josimarz/gopher-pacman/internal/game/stack"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

type ghostReachedTileEvent struct {
	point     *tile.Point
	timestamp time.Time
}

func newGhostReachedTileEvent(pt *tile.Point) *ghostReachedTileEvent {
	return &ghostReachedTileEvent{
		point:     pt,
		timestamp: time.Now(),
	}
}

func (e *ghostReachedTileEvent) GetName() string {
	return "ghost.reached.tile"
}

func (e *ghostReachedTileEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *ghostReachedTileEvent) GetPayload() any {
	return e.point
}

type GhostTracking struct {
	path      *stack.Stack[tile.Point]
	dir       direction.Direction
	prevPoint *tile.Point
	currPoint *tile.Point
	nextPoint *tile.Point
	speed     int
}

func NewGhostTracking(pt *tile.Point) *GhostTracking {
	return &GhostTracking{
		dir:       direction.Up,
		currPoint: pt,
		nextPoint: pt,
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
		e := newGhostReachedTileEvent(t.currPoint)
		event.Dispatcher().Dispatch(e)
		t.nextDir()
	}
	t.moveX()
	t.moveY()
}

func (t *GhostTracking) nextDir() {
	for {
		if dir := direction.Direction(rand.Intn(4)); t.goNext(dir) {
			t.dir = dir
			return
		}
	}
}

func (t *GhostTracking) goNext(dir direction.Direction) bool {
	switch dir {
	case direction.Up:
		t.nextPoint.Y -= tile.Size
	case direction.Down:
		t.nextPoint.Y += tile.Size
	case direction.Left:
		t.nextPoint.X -= tile.Size
	case direction.Right:
		t.nextPoint.X += tile.Size
	}
	if !world.Instance().Accessible(t.nextPoint) {
		t.nextPoint = t.currPoint.Clone()
		return false
	}
	return true
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
