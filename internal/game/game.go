package game

import (
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/event"
	"github.com/josimarz/gopher-pacman/internal/game/ghost"
	_ "github.com/josimarz/gopher-pacman/internal/game/handler"
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

type GameStartedEvent struct {
	timestamp time.Time
}

func NewGameStartedEvent() *GameStartedEvent {
	return &GameStartedEvent{}
}

func (e *GameStartedEvent) GetName() string {
	return "game.started"
}

func (e *GameStartedEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *GameStartedEvent) GetPayload() any {
	return struct{}{}
}

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
		e := NewGameStartedEvent()
		event.Dispatcher().Dispatch(e)
	})
	return game
}

func (g *Game) Update() error {
	input.Instance().Listen()
	player.Instance().Update()
	ghost.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	world.Instance().Draw(screen)
	player.Instance().Draw(screen)
	ghost.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
