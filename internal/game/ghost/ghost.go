package ghost

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/assets"
	"github.com/josimarz/gopher-pacman/internal/game/move"
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
)

type Color uint8

const (
	Red Color = iota
	Pink
	Cyan
	Orange
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

type Ghost struct {
	color    Color
	tracking *move.GhostTracking
}

func new(color Color) *Ghost {
	return &Ghost{
		color:    color,
		tracking: move.NewGhostTracking(startPoint(color)),
	}
}

func startPoint(color Color) *point.Point {
	switch color {
	case Red:
		return point.New(10*tile.Size, 7*tile.Size)
	case Pink:
		return point.New(10*tile.Size, 9*tile.Size)
	case Cyan:
		return point.New(9*tile.Size, 9*tile.Size)
	case Orange:
		return point.New(11*tile.Size, 9*tile.Size)
	default:
		log.Fatalf("Invalid color: %d", color)
		return nil
	}
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
	x := int(g.color)
	y := int(g.tracking.Dir()) + 2
	return x, y
}
