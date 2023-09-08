package ghost

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/assets"
	"github.com/josimarz/gopher-pacman/internal/game/move"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
)

type Color uint8
type FearStatus uint8

const (
	Red Color = iota
	Pink
	Cyan
	Orange
)

const (
	None FearStatus = iota
	Frightened
	Recovering
)

var (
	Blinky, Pinky, Inky, Clyde *Ghost
)

func init() {
	Blinky = new(Red)
	Pinky = new(Pink)
	Inky = new(Cyan)
	Clyde = new(Orange)
}

func Update() {
	Blinky.update()
	Pinky.update()
	Inky.update()
	Clyde.update()
}

func Draw(screen *ebiten.Image) {
	Blinky.draw(screen)
	Pinky.draw(screen)
	Inky.draw(screen)
	Clyde.draw(screen)
}

func Frighten() {
	go Blinky.frighten()
	go Pinky.frighten()
	go Inky.frighten()
	go Clyde.frighten()
}

func CheckCollisions(pt *tile.Point) []*Ghost {
	ghosts := []*Ghost{}
	if Blinky.currPoint().Collide(pt) {
		ghosts = append(ghosts, Blinky)
	}
	if Pinky.currPoint().Collide(pt) {
		ghosts = append(ghosts, Pinky)
	}
	if Inky.currPoint().Collide(pt) {
		ghosts = append(ghosts, Inky)
	}
	if Clyde.currPoint().Collide(pt) {
		ghosts = append(ghosts, Clyde)
	}
	return ghosts
}

type Ghost struct {
	dead       bool
	color      Color
	fearStatus FearStatus
	tracking   *move.GhostTracking
}

func new(color Color) *Ghost {
	return &Ghost{
		color:    color,
		tracking: move.NewGhostTracking(startPoint(color)),
	}
}

func startPoint(color Color) *tile.Point {
	switch color {
	case Red:
		return tile.NewPoint(10*tile.Size, 7*tile.Size)
	case Pink:
		return tile.NewPoint(10*tile.Size, 9*tile.Size)
	case Cyan:
		return tile.NewPoint(9*tile.Size, 9*tile.Size)
	case Orange:
		return tile.NewPoint(11*tile.Size, 9*tile.Size)
	default:
		log.Fatalf("Invalid color: %d", color)
		return nil
	}
}

func (g *Ghost) FearStatus() FearStatus {
	return g.fearStatus
}

func (g *Ghost) Die() {
	g.dead = true
	g.fearStatus = None
}

func (g *Ghost) currPoint() *tile.Point {
	return g.tracking.CurrPoint()
}

func (g *Ghost) update() {
	g.tracking.Move()
}

func (g *Ghost) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.tracking.CurrPoint().X), float64(g.tracking.CurrPoint().Y))
	x, y := g.spriteCoords()
	sx, sy := x*tile.Size, y*tile.Size
	r := image.Rect(sx, sy, sx+tile.Size, sy+tile.Size)
	img := assets.SpriteSheet.SubImage(r).(*ebiten.Image)
	screen.DrawImage(img, op)
}

func (g *Ghost) spriteCoords() (int, int) {
	if g.fearStatus == Frightened {
		return 5, 2
	}
	if g.fearStatus == Recovering {
		return 5, 3
	}
	x := int(g.color)
	if g.dead {
		x = 4
	}
	y := int(g.tracking.Dir()) + 2
	return x, y
}

func (g *Ghost) frighten() {
	g.fearStatus = Frightened
	time.Sleep(5 * time.Second)
	g.fearStatus = Recovering
	time.Sleep(5 * time.Second)
	g.fearStatus = None
}
