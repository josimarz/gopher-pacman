package player

import (
	"image"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/assets"
	"github.com/josimarz/gopher-pacman/internal/game/direction"
	"github.com/josimarz/gopher-pacman/internal/game/move"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
)

var (
	once     sync.Once
	instance *Player
)

type Player struct {
	mouth    *mouth
	tracking *move.PlayerTracking
}

func Instance() *Player {
	once.Do(func() {
		instance = &Player{
			mouth:    newMouth(),
			tracking: move.NewPlayerTracking(),
		}
	})
	return instance
}

func (p *Player) ChangeDir(dir direction.Direction) {
	p.tracking.ChangeDir(dir)
}

func (p *Player) Update() {
	p.mouth.update()
	p.tracking.Move()
}

func (p *Player) PrevPoint() *tile.Point {
	return p.tracking.PrevPoint()
}

func (p *Player) CurrPoint() *tile.Point {
	return p.tracking.CurrPoint()
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.tracking.CurrPoint().X), float64(p.tracking.CurrPoint().Y))
	x, y := p.spriteCoords()
	sx, sy := x*tile.Size, y*tile.Size
	r := image.Rect(sx, sy, sx+tile.Size, sy+tile.Size)
	img := assets.SpriteSheet.SubImage(r).(*ebiten.Image)
	screen.DrawImage(img, op)
}

func (p *Player) spriteCoords() (int, int) {
	if p.mouth.status == mouthClosed {
		return 4, 0
	}
	x := 0
	y := 0
	switch p.tracking.CurrDir() {
	case direction.Up:
		x = 3
	case direction.Down:
		x = 2
	case direction.Left:
		x = 0
	case direction.Right:
		x = 1
	}
	if p.mouth.status == mouthOpen {
		y = 1
	}
	return x, y
}
