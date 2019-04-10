package gcounter

import (
	"context"
)

type CounterMessage interface {
	handle(c *Counter)
}

type Counter struct {
	value int
	inbox chan CounterMessage
}

func (c *Counter) Run(ctx context.Context) {
	c.value = 0
	c.inbox = make(chan CounterMessage, 100)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-c.inbox:
				msg.handle(c)
			}
		}
	}()
}

func (c Counter) Increment() {
	c.inbox <- &Increment{}
}

func (c Counter) Decrement() {
	c.inbox <- &Decrement{}
}

func (c Counter) GetValue() int {
	valueOut := make(chan int)
	c.inbox <- &GetValue{valueOut}

	return <-valueOut
}
