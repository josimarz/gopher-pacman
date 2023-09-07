package event

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josimarz/gopher-pacman/internal/game/player"
)

func init() {
	Dispatcher().Attach("key.pressed", HandleKeyPressed)
}

func HandleKeyPressed(e Event) {
	var dir player.Direction
	key := e.GetPayload().(ebiten.Key)
	switch key.String() {
	case "ArrowUp", "W":
		dir = player.Up
	case "ArrowDown", "S":
		dir = player.Down
	case "ArrowLeft", "A":
		dir = player.Left
	case "ArrowRight", "D":
		dir = player.Right
	}
	player.Instance().ChangeDirection(dir)
}
