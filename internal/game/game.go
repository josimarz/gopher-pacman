package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/josimarz/gopher-pacman/internal/game/direction"
	"github.com/josimarz/gopher-pacman/internal/game/ghost"
	"github.com/josimarz/gopher-pacman/internal/game/player"
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/sfx"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

const (
	screenWidth  = world.Cols * tile.Size
	screenHeight = world.Rows * tile.Size
)

type Game struct{}

func Start() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pacman")
	sfx.PlayGameStart()
	if err := ebiten.RunGame(New()); err != nil {
		log.Fatal(err)
	}
}

func New() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	g.handleInput()
	g.checkCollisions()
	player.Update()
	ghost.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	world.Draw(screen)
	player.Draw(screen)
	ghost.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) handleInput() {
	keys := inpututil.AppendPressedKeys(nil)
	for _, key := range keys {
		var dir direction.Direction
		switch key.String() {
		case "ArrowUp", "W":
			dir = direction.Up
		case "ArrowDown", "S":
			dir = direction.Down
		case "ArrowLeft", "A":
			dir = direction.Left
		case "ArrowRight", "D":
			dir = direction.Right
		default:
			continue
		}
		player.SetNextDir(dir)
	}
}

func (g *Game) checkCollisions() {
	ghosts := []*ghost.Ghost{
		ghost.Blinky,
		ghost.Pinky,
		ghost.Inky,
		ghost.Clyde,
	}
	for _, gh := range ghosts {
		collide := g.checkCollision(player.CurrPoint(), gh.CurrPoint())
		if collide {
			if gh.Status() == ghost.Alive {
				player.Die()
			} else if gh.Status() != ghost.Dead {
				gh.Die()
				sfx.PlayEatGhost()
			}
		}
	}
}

func (g *Game) checkCollision(p1, p2 *point.Point) bool {
	return (p1.X <= p2.X && p1.X+tile.Size >= p2.X && p1.Y <= p2.Y && p1.Y+tile.Size >= p2.Y) ||
		(p2.X <= p1.X && p2.X+tile.Size >= p1.X && p2.Y <= p1.Y && p2.Y+tile.Size >= p1.Y)
}
