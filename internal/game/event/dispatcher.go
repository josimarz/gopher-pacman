package event

import (
	"sync"
)

var (
	once     sync.Once
	instance *dispatcher
)

type dispatcher struct {
	handlers map[string][]EventHandler
}

func Dispatcher() *dispatcher {
	once.Do(func() {
		instance = &dispatcher{
			handlers: make(map[string][]EventHandler),
		}
	})
	return instance
}

func (d *dispatcher) Dispatch(e Event) {
	if handlers, ok := d.handlers[e.GetName()]; ok {
		for _, handler := range handlers {
			handler(e)
		}
	}
}

func (d *dispatcher) Attach(name string, handler EventHandler) *dispatcher {
	d.handlers[name] = append(d.handlers[name], handler)
	return d
}

func (d *dispatcher) Clear() {
	d.handlers = make(map[string][]EventHandler)
}
