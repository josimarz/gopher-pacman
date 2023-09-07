package input

import (
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/josimarz/gopher-pacman/internal/game/event"
)

var (
	once     sync.Once
	listener *InputListener
)

type KeyPressedEvent struct {
	key       ebiten.Key
	timestamp time.Time
}

func NewKeyPressedEvent(key ebiten.Key) *KeyPressedEvent {
	return &KeyPressedEvent{
		key:       key,
		timestamp: time.Now(),
	}
}

func (e *KeyPressedEvent) GetName() string {
	return "key.pressed"
}

func (e *KeyPressedEvent) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *KeyPressedEvent) GetPayload() any {
	return e.key
}

type InputListener struct {
}

func Instance() *InputListener {
	once.Do(func() {
		listener = &InputListener{}
	})
	return listener
}

func (l *InputListener) Listen() {
	keys := inpututil.AppendPressedKeys(nil)
	for _, key := range keys {
		e := NewKeyPressedEvent(key)
		event.Dispatcher().Dispatch(e)
	}
}
