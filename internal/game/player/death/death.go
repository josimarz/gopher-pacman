package death

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/assets"
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
)

type Animation struct {
	delta float64
	speed float64
	point *point.Point
}

func NewAnimation() *Animation {
	return &Animation{
		delta: 1,
		speed: 0.03,
	}
}

func (a *Animation) Start(point *point.Point) {
	a.delta = 1
	a.point = point
}

func (a *Animation) Draw(screen *ebiten.Image) bool {
	a.delta += a.speed * a.delta
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(a.point.X), float64(a.point.Y))
	x, y := int(a.delta), 6
	sx, sy := x*tile.Size, y*tile.Size
	r := image.Rect(sx, sy, sx+tile.Size, sy+tile.Size)
	img := assets.SpriteSheet.SubImage(r).(*ebiten.Image)
	screen.DrawImage(img, op)
	return a.delta >= 11
}
