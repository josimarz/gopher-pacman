package tile

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/josimarz/gopher-pacman/internal/game/position"
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

type Tile struct {
	content Content
	pos     *position.Position
}

func New(content Content, pos *position.Position) *Tile {
	return &Tile{
		content: content,
		pos:     pos,
	}
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
	x := float32(t.pos.X * Size)
	y := float32(t.pos.Y * Size)
	clr := color.RGBA{
		R: 60,
		G: 94,
		B: 164,
	}
	vector.DrawFilledRect(screen, x, y, float32(Size), float32(Size), clr, true)
}

func (t *Tile) drawDoor(screen *ebiten.Image) {
	x := float32(t.pos.X * Size)
	y := float32(t.pos.Y * Size)
	clr := color.RGBA{
		R: 141,
		G: 141,
		B: 141,
	}
	vector.DrawFilledRect(screen, x, y, float32(Size), float32(Size), clr, true)
}

func (t *Tile) drawDot(screen *ebiten.Image) {
	cx := float32(t.pos.X*Size + Size/2)
	cy := float32(t.pos.Y*Size + Size/2)
	r := float32(5)
	vector.DrawFilledCircle(screen, cx, cy, r, color.White, true)
}

func (t *Tile) drawPill(screen *ebiten.Image) {
	cx := float32(t.pos.X*Size + Size/2)
	cy := float32(t.pos.Y*Size + Size/2)
	r := float32(10)
	vector.DrawFilledCircle(screen, cx, cy, r, color.White, true)
}
