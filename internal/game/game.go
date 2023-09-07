package game

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/josimarz/gopher-pacman/internal/game/event"
	"github.com/josimarz/gopher-pacman/internal/game/input"
	"github.com/josimarz/gopher-pacman/internal/game/player"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

const (
	Width  = tile.Size * world.Cols
	Height = tile.Size * world.Rows
)

var (
	once sync.Once
	game *Game
)

type Game struct{}

func Start() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Pacman")
	if err := ebiten.RunGame(Instance()); err != nil {
		log.Fatal(err)
	}
}

func Instance() *Game {
	once.Do(func() {
		game = &Game{}
	})
	return game
}

func (g *Game) Update() error {
	input.Instance().Listen()
	player.Instance().Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	world.Instance().Draw(screen)
	player.Instance().Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
