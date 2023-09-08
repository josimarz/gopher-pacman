package move

import (
	"time"

	"github.com/josimarz/gopher-pacman/internal/game/event"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

type playerReachedTileEvent struct {
	point     *tile.Point
	timestamp time.Time
}

func newPlayerReachedTileEvent(point *tile.Point) *playerReachedTileEvent {
	return &playerReachedTileEvent{
		point:     point,
		timestamp: time.Now(),
	}
}

func (e *playerReachedTileEvent) GetName() string {
	return "player.reached.tile"
}

func (e *playerReachedTileEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *playerReachedTileEvent) GetPayload() any {
	return e.point
}

type PlayerTracking struct {
	currDir   Direction
	nextDir   Direction
	currPoint *tile.Point
	nextPoint *tile.Point
	speed     int
}

func NewPlayerTracking() *PlayerTracking {
	return &PlayerTracking{
		currDir:   Up,
		nextDir:   Up,
		currPoint: tile.NewPoint(10*tile.Size, 15*tile.Size),
		nextPoint: tile.NewPoint(10*tile.Size, 15*tile.Size),
		speed:     2,
	}
}

func (t *PlayerTracking) CurrPoint() *tile.Point {
	return t.currPoint
}

func (t *PlayerTracking) CurrDir() Direction {
	return t.currDir
}

func (t *PlayerTracking) ChangeDir(dir Direction) {
	if t.nextDir != dir {
		t.nextDir = dir
	}
}

func (t *PlayerTracking) Move() {
	if t.currPoint.Equals(t.nextPoint) {
		e := newPlayerReachedTileEvent(t.nextPoint)
		event.Dispatcher().Dispatch(e)
		if t.goNext(t.nextDir) {
			t.currDir = t.nextDir
		} else if !t.goNext(t.currDir) {
			return
		}
	}
	t.moveX()
	t.moveY()
}

func (t *PlayerTracking) goNext(dir Direction) bool {
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

func (t *PlayerTracking) moveX() {
	if t.currPoint.X > t.nextPoint.X {
		t.currPoint.X -= t.speed
	} else if t.currPoint.X < t.nextPoint.X {
		t.currPoint.X += t.speed
	}
}

func (t *PlayerTracking) moveY() {
	if t.currPoint.Y > t.nextPoint.Y {
		t.currPoint.Y -= t.speed
	} else if t.currPoint.Y < t.nextPoint.Y {
		t.currPoint.Y += t.speed
	}
}
