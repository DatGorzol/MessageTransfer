package websocket

import "log"

type Dispatcher struct {
	queue chan Envelope
}

func NewDispatcher(buffer int) *Dispatcher {
	return &Dispatcher{
		queue: make(chan Envelope, buffer),
	}
}

func (d *Dispatcher) Start() {
	go func() {
		for env := range d.queue {
			target, ok := manager.Get(env.To)
			if !ok || !target.Auth {
				log.Printf("dispatcher: target %s not available\n", env.To)
				continue
			}

			if err := target.Send(env.Data); err != nil {
				log.Println("dispatcher send error:", err)
			}
		}
	}()
}

func (d *Dispatcher) Dispatch(env Envelope) {
	d.queue <- env
}

var dispatcher = NewDispatcher(1000)

func StartDispatcher() {
	dispatcher.Start()
}
