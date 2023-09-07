package move

import (
	"math/rand"

	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

type GhostTracking struct {
	dir       Direction
	currPoint *point.Point
	nextPoint *point.Point
	speed     int
}

func NewGhostTracking(pt *point.Point) *GhostTracking {
	return &GhostTracking{
		dir:       Up,
		currPoint: pt,
		nextPoint: pt,
		speed:     1,
	}
}

func (t *GhostTracking) CurrPoint() *point.Point {
	return t.currPoint
}

func (t *GhostTracking) Dir() Direction {
	return t.dir
}

func (t *GhostTracking) Move() {
	if t.currPoint.Equals(t.nextPoint) {
		t.nextDir()
	}
	t.moveX()
	t.moveY()
}

func (t *GhostTracking) nextDir() {
	for {
		if dir := Direction(rand.Intn(4)); t.goNext(dir) {
			t.dir = dir
			return
		}
	}
}

func (t *GhostTracking) goNext(dir Direction) bool {
	switch dir {
	case Up:
		t.nextPoint.Y -= tile.Size
	case Down:
		t.nextPoint.Y += tile.Size
	case Left:
		t.nextPoint.X -= tile.Size
	case Right:
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
