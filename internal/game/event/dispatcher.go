package event

import (
	"sync"
)

var (
	once       sync.Once
	dispatcher *EventDispatcher
)

type EventDispatcher struct {
	handlers map[string][]EventHandler
}

func Dispatcher() *EventDispatcher {
	once.Do(func() {
		dispatcher = &EventDispatcher{
			handlers: make(map[string][]EventHandler),
		}
	})
	return dispatcher
}

func (d *EventDispatcher) Dispatch(e Event) {
	if handlers, ok := d.handlers[e.GetName()]; ok {
		for _, handler := range handlers {
			handler(e)
		}
	}
}

func (d *EventDispatcher) Attach(name string, handler EventHandler) *EventDispatcher {
	d.handlers[name] = append(d.handlers[name], handler)
	return d
}

func (d *EventDispatcher) Clear() {
	d.handlers = make(map[string][]EventHandler)
}
