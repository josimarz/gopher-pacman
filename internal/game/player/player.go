package player

import (
	"image"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/assets"
	"github.com/josimarz/gopher-pacman/internal/game/position"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
)

type mouthStatus uint8
type Direction uint8

const (
	mouthOpen mouthStatus = iota + 1
	mouthClosing
	mouthClosed
	mouthOpening
)

const (
	Up Direction = iota
	Down
	Left
	Right
)

var (
	once   sync.Once
	player *Player
)

type Mouth struct {
	status mouthStatus
	delta  float64
	speed  float64
}

func (m *Mouth) Update() {
	m.delta *= m.speed
	if m.delta >= 5 {
		m.delta = 1
	}
	m.status = mouthStatus(m.delta)
}

type Player struct {
	pos   *position.Position
	dir   Direction
	mouth *Mouth
	speed uint8
}

func Instance() *Player {
	once.Do(func() {
		player = &Player{
			pos:   position.New(10, 15),
			dir:   Up,
			speed: 1,
			mouth: &Mouth{
				status: mouthClosed,
				delta:  1,
				speed:  1.05,
			},
		}
	})
	return player
}

func (p *Player) ChangeDirection(dir Direction) {
	p.dir = dir
}

func (p *Player) Update() {
	p.mouth.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.pos.X*tile.Size), float64(p.pos.Y*tile.Size))
	x, y := p.spriteCoords()
	sx, sy := x*tile.Size, y*tile.Size
	screen.DrawImage(assets.SpriteSheet.SubImage(image.Rect(sx, sy, sx+tile.Size, sy+tile.Size)).(*ebiten.Image), op)
}

func (p *Player) spriteCoords() (int, int) {
	if p.mouth.status == mouthClosed {
		return 4, 0
	}
	x := 0
	y := 0
	switch p.dir {
	case Up:
		x = 3
	case Down:
		x = 2
	case Left:
		x = 0
	case Right:
		x = 1
	}
	if p.mouth.status == mouthOpen {
		y = 1
	}
	return x, y
}
