package player

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/assets"
	"github.com/josimarz/gopher-pacman/internal/game/direction"
	"github.com/josimarz/gopher-pacman/internal/game/ghost"
	"github.com/josimarz/gopher-pacman/internal/game/player/death"
	"github.com/josimarz/gopher-pacman/internal/game/player/mouth"
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/sfx"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

var (
	player     *Player
	startPoint = point.New(10*tile.Size, 15*tile.Size)
)

func init() {
	player = &Player{
		mouth:     mouth.New(),
		currDir:   direction.Up,
		nextDir:   direction.Up,
		currPoint: startPoint.Clone(),
		nextPoint: startPoint.Clone(),
		speed:     2,
		deathAnim: death.NewAnimation(),
	}
}

func Update() {
	player.update()
}

func Draw(screen *ebiten.Image) {
	player.draw(screen)
}

func CurrPoint() *point.Point {
	return player.currPoint
}

func SetNextDir(dir direction.Direction) {
	player.nextDir = dir
}

func Die() {
	player.die()
}

type Player struct {
	dead      bool
	mouth     *mouth.Mouth
	currDir   direction.Direction
	nextDir   direction.Direction
	currPoint *point.Point
	nextPoint *point.Point
	speed     int
	deathAnim *death.Animation
}

func (p *Player) update() {
	if !p.dead {
		p.mouth.Update()
		p.Move()
	}
}

func (p *Player) setNextDir(dir direction.Direction) {
	p.nextDir = dir
}

func (p *Player) die() {
	if !p.dead {
		p.dead = true
		p.deathAnim.Start(p.currPoint)
		sfx.PlayDeath()
	}
}

func (p *Player) draw(screen *ebiten.Image) {
	if p.dead {
		done := p.deathAnim.Draw(screen)
		if done {
			go func() {
				time.Sleep(500 * time.Millisecond)
				p.respawn()
			}()
		}
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(p.currPoint.X), float64(p.currPoint.Y))
		x, y := p.spriteCoords()
		sx, sy := x*tile.Size, y*tile.Size
		r := image.Rect(sx, sy, sx+tile.Size, sy+tile.Size)
		img := assets.SpriteSheet.SubImage(r).(*ebiten.Image)
		screen.DrawImage(img, op)
	}
}

func (p *Player) respawn() {
	p.currDir = direction.Up
	p.nextDir = direction.Down
	p.currPoint = startPoint.Clone()
	p.nextPoint = startPoint.Clone()
	p.dead = false
}

func (p *Player) spriteCoords() (int, int) {
	if p.mouth.Status() == mouth.Closed {
		return 4, 0
	}
	dirs := []int{3, 2, 0, 1}
	x := dirs[int(p.currDir)]
	y := 0
	if p.mouth.Status() == mouth.Open {
		y = 1
	}
	return x, y
}

func (p *Player) Move() {
	if p.currPoint.Equals(p.nextPoint) {
		p.tryEat()
		if p.changeDir() {
			p.currDir = p.nextDir
		} else {
			p.nextDir = p.currDir
			if !p.changeDir() {
				return
			}
		}
	}
	p.moveX()
	p.moveY()
}

func (p *Player) tryEat() {
	t := world.TileAt(p.currPoint)
	if t != nil {
		if t.Content() == tile.Dot {
			t.RemoveContent()
			sfx.PlayMunch1()
			sfx.PlayMunch2()
		} else if t.Content() == tile.PowerPellet {
			t.RemoveContent()
			sfx.PlayPowerPellet()
			ghost.DizzyAll()
		}
	}
}

func (p *Player) changeDir() bool {
	pt := p.nextPoint.Clone()
	switch p.nextDir {
	case direction.Up:
		pt.Y -= tile.Size
	case direction.Down:
		pt.Y += tile.Size
	case direction.Left:
		pt.X -= tile.Size
	case direction.Right:
		pt.X += tile.Size
	}
	if world.Reachable(pt) {
		p.nextPoint = pt
		return true
	}
	return false
}

func (p *Player) moveX() {
	if p.currPoint.X > p.nextPoint.X {
		p.currPoint.X -= p.speed
	} else if p.currPoint.X < p.nextPoint.X {
		p.currPoint.X += p.speed
	}
}

func (p *Player) moveY() {
	if p.currPoint.Y > p.nextPoint.Y {
		p.currPoint.Y -= p.speed
	} else if p.currPoint.Y < p.nextPoint.Y {
		p.currPoint.Y += p.speed
	}
}
