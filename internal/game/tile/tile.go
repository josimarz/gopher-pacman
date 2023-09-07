package tile

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/josimarz/gopher-pacman/internal/game/event"
	"github.com/josimarz/gopher-pacman/internal/game/point"
)

type Content uint8

const (
	Size int = 32
)

const (
	None Content = iota
	Wall
	Door
	Dot
	Pill
)

type DotEatenEvent struct {
	timestamp time.Time
}

func NewDotEatenEvent() *DotEatenEvent {
	return &DotEatenEvent{
		timestamp: time.Now(),
	}
}

func (e *DotEatenEvent) GetName() string {
	return "dot.eaten"
}

func (e *DotEatenEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *DotEatenEvent) GetPayload() any {
	return struct{}{}
}

type PillEatenEvent struct {
	timestamp time.Time
}

func NewPillEatenEvent() *PillEatenEvent {
	return &PillEatenEvent{
		timestamp: time.Now(),
	}
}

func (e *PillEatenEvent) GetName() string {
	return "pill.eaten"
}

func (e *PillEatenEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *PillEatenEvent) GetPayload() any {
	return struct{}{}
}

type Tile struct {
	content Content
	point   *point.Point
}

func New(content Content, pt *point.Point) *Tile {
	return &Tile{
		content: content,
		point:   pt,
	}
}

func (t *Tile) Eat() {
	if t.content == Dot {
		e := NewDotEatenEvent()
		event.Dispatcher().Dispatch(e)
		t.content = None
	}
	if t.content == Pill {
		e := NewPillEatenEvent()
		event.Dispatcher().Dispatch(e)
		t.content = None
	}
}

func (t *Tile) Accessible() bool {
	return t.content == None || t.content == Dot || t.content == Pill
}

func (t *Tile) Point() *point.Point {
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
