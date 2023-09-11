package tile

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/josimarz/gopher-pacman/internal/game/event"
)

type Content uint8

const (
	Size int = 32
)

const (
	None Content = iota
	Outline
	Wall
	Door
	Dot
	Pill
)

type dotEatenEvent struct {
	timestamp time.Time
}

func newDotEatenEvent() *dotEatenEvent {
	return &dotEatenEvent{
		timestamp: time.Now(),
	}
}

func (e *dotEatenEvent) GetName() string {
	return "dot.eaten"
}

func (e *dotEatenEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *dotEatenEvent) GetPayload() any {
	return struct{}{}
}

type pillEatenEvent struct {
	timestamp time.Time
}

func newPillEatenEvent() *pillEatenEvent {
	return &pillEatenEvent{
		timestamp: time.Now(),
	}
}

func (e *pillEatenEvent) GetName() string {
	return "pill.eaten"
}

func (e *pillEatenEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *pillEatenEvent) GetPayload() any {
	return struct{}{}
}

type Tile struct {
	content Content
	point   *Point
}

func New(content Content, pt *Point) *Tile {
	return &Tile{
		content: content,
		point:   pt,
	}
}

func (t *Tile) Eat() {
	if t.content == Dot {
		e := newDotEatenEvent()
		event.Dispatcher().Dispatch(e)
		t.content = None
	}
	if t.content == Pill {
		e := newPillEatenEvent()
		event.Dispatcher().Dispatch(e)
		t.content = None
	}
}

func (t *Tile) Accessible() bool {
	return t.content != Wall
}

func (t *Tile) Point() *Point {
	return t.point
}

func (t *Tile) Draw(screen *ebiten.Image) {
	switch t.content {
	case Wall:
		t.drawWall(screen)
	case Door:
		t.drawDoor(screen)
	case Dot:
		t.drawDot(screen)
	case Pill:
		t.drawPill(screen)
	}
}

func (t *Tile) drawWall(screen *ebiten.Image) {
	x := float32(t.point.X)
	y := float32(t.point.Y)
	clr := color.RGBA{
		R: 60,
		G: 94,
		B: 164,
	}
	vector.DrawFilledRect(screen, x, y, float32(Size), float32(Size), clr, true)
}

func (t *Tile) drawDoor(screen *ebiten.Image) {
	x := float32(t.point.X)
	y := float32(t.point.Y)
	clr := color.RGBA{
		R: 141,
		G: 141,
		B: 141,
	}
	vector.DrawFilledRect(screen, x, y, float32(Size), float32(Size), clr, true)
}

func (t *Tile) drawDot(screen *ebiten.Image) {
	cx := float32(t.point.X + Size/2)
	cy := float32(t.point.Y + Size/2)
	r := float32(5)
	vector.DrawFilledCircle(screen, cx, cy, r, color.White, true)
}

func (t *Tile) drawPill(screen *ebiten.Image) {
	cx := float32(t.point.X + Size/2)
	cy := float32(t.point.Y + Size/2)
	r := float32(10)
	vector.DrawFilledCircle(screen, cx, cy, r, color.White, true)
}
