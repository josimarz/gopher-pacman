package handler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/direction"
	"github.com/josimarz/gopher-pacman/internal/game/event"
	"github.com/josimarz/gopher-pacman/internal/game/ghost"
	"github.com/josimarz/gopher-pacman/internal/game/player"
	"github.com/josimarz/gopher-pacman/internal/game/sfx"
	"github.com/josimarz/gopher-pacman/internal/game/tile"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

func init() {
	event.Dispatcher().
		Attach("game.started", handleGameStarted).
		Attach("key.pressed", handleKeyPressed).
		Attach("player.reached.tile", handlePlayerReachedTile).
		Attach("dot.eaten", handleDotEaten).
		Attach("pill.eaten", handlePillEaten).
		Attach("ghost.died", handleGhostDied)
}

func handleGameStarted(e event.Event) {
	go sfx.AudioPlayer().PlayGameStart()
}

func handleKeyPressed(e event.Event) {
	if key, ok := e.GetPayload().(ebiten.Key); ok {
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
		}
		player.Instance().ChangeDir(dir)
	}
}

func handlePlayerReachedTile(e event.Event) {
	if pt, ok := e.GetPayload().(*tile.Point); ok {
		ghosts := ghost.CheckCollisions(pt)
		for _, g := range ghosts {
			if g.FearStatus() != ghost.None {
				g.Die()
			}
		}
		tile := world.Instance().TileAt(pt)
		if tile != nil {
			tile.Eat()
		}
	}
}

func handleDotEaten(e event.Event) {
	go func() {
		sfx.AudioPlayer().PlayMunch1()
		sfx.AudioPlayer().PlayMunch2()
	}()
}

func handlePillEaten(e event.Event) {
	go sfx.AudioPlayer().PlayPowerPellet()
	go ghost.Frighten()
}

func handleGhostDied(e event.Event) {
	go sfx.AudioPlayer().PlayEatGhost()
}
