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
	width  = tile.Size * world.Cols
	height = tile.Size * world.Rows
)

var (
	once     sync.Once
	instance *Game
)

type gameStartedEvent struct {
	timestamp time.Time
}

func newGameStartedEvent() *gameStartedEvent {
	return &gameStartedEvent{}
}

func (e *gameStartedEvent) GetName() string {
	return "game.started"
}

func (e *gameStartedEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *gameStartedEvent) GetPayload() any {
	return struct{}{}
}

type Game struct{}

func Start() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Pacman")
	if err := ebiten.RunGame(Instance()); err != nil {
		log.Fatal(err)
	}
}

func Instance() *Game {
	once.Do(func() {
		instance = &Game{}
		e := newGameStartedEvent()
		event.Dispatcher().Dispatch(e)
	})
	return instance
}

func (g *Game) Update() error {
	go input.Listen()
	go player.Instance().Update()
	go ghost.Update()
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
