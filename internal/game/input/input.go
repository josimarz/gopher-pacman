package input

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/josimarz/gopher-pacman/internal/game/event"
)

type keyPressedEvent struct {
	key       ebiten.Key
	timestamp time.Time
}

func newKeyPressedEvent(key ebiten.Key) *keyPressedEvent {
	return &keyPressedEvent{
		key:       key,
		timestamp: time.Now(),
	}
}

func (e *keyPressedEvent) GetName() string {
	return "key.pressed"
}

func (e *keyPressedEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *keyPressedEvent) GetPayload() any {
	return e.key
}

func Listen() {
	keys := inpututil.AppendPressedKeys(nil)
	for _, key := range keys {
		e := newKeyPressedEvent(key)
		event.Dispatcher().Dispatch(e)
	}
}
