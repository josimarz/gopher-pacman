package handler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/event"
	"github.com/josimarz/gopher-pacman/internal/game/move"
	"github.com/josimarz/gopher-pacman/internal/game/player"
	"github.com/josimarz/gopher-pacman/internal/game/point"
	"github.com/josimarz/gopher-pacman/internal/game/sfx"
	"github.com/josimarz/gopher-pacman/internal/game/world"
)

func init() {
	event.Dispatcher().
		Attach("game.started", HandleGameStarted).
		Attach("key.pressed", HandleKeyPressed).
		Attach("player.reached.tile", HandlePlayerReachedTile).
		Attach("dot.eaten", HandleDotEaten).
		Attach("pill.eaten", HandlePillEaten)
}

func HandleGameStarted(e event.Event) {
	sfx.AudioPlayer().PlayGameStart()
}

func HandleKeyPressed(e event.Event) {
	if key, ok := e.GetPayload().(ebiten.Key); ok {
		var dir move.Direction
		switch key.String() {
		case "ArrowUp", "W":
			dir = move.Up
		case "ArrowDown", "S":
			dir = move.Down
		case "ArrowLeft", "A":
			dir = move.Left
		case "ArrowRight", "D":
			dir = move.Right
		}
		player.Instance().ChangeDir(dir)
	}
}

func HandlePlayerReachedTile(e event.Event) {
	if pt, ok := e.GetPayload().(*point.Point); ok {
		tile := world.Instance().TileAt(pt)
		if tile != nil {
			tile.Eat()
		}
	}
}

func HandleDotEaten(e event.Event) {
	sfx.AudioPlayer().PlayMunch1()
	sfx.AudioPlayer().PlayMunch2()
}

func HandlePillEaten(e event.Event) {
	sfx.AudioPlayer().PlayPowerPellet()
}
