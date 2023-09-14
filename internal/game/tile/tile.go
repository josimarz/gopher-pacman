package tile

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/josimarz/gopher-pacman/internal/game/point"
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
	PowerPellet
)

type Tile struct {
	content Content
	point   *point.Point
}

func New(content Content, point *point.Point) *Tile {
	return &Tile{
		content: content,
		point:   point,
	}
}

func (t *Tile) Point() *point.Point {
	return t.point
}

func (t *Tile) Reachable() bool {
	return t.content == None || t.content == Door || t.content == Dot || t.content == PowerPellet
}

func (t *Tile) Draw(screen *ebiten.Image) {
	switch t.content {
	case Wall:
		t.drawWall(screen)
	case Door:
		t.drawDoor(screen)
	case Dot:
		t.drawDot(screen)
	case PowerPellet:
		t.drawPill(screen)
	}
}

func (t *Tile) Content() Content {
	return t.content
}

func (t *Tile) RemoveContent() {
	t.content = None
}

func (t *Tile) drawWall(screen *ebiten.Image) {
	x := float32(t.point.X)
	y := float32(t.point.Y)
	w := float32(Size)
	h := float32(Size)
	clr := color.RGBA{
		R: 60,
		G: 94,
		B: 164,
	}
	vector.DrawFilledRect(screen, x, y, w, h, clr, true)
}

func (t *Tile) drawDoor(screen *ebiten.Image) {
	x := float32(t.point.X)
	y := float32(t.point.Y)
	w := float32(Size)
	h := float32(Size)
	clr := color.RGBA{
		R: 141,
		G: 141,
		B: 141,
	}
	vector.DrawFilledRect(screen, x, y, w, h, clr, true)
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
