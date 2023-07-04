package main

import (
	"fmt"
	"time"
)

type Event struct {
	ID        int
	Timestamp time.Time
}

type EventHandler interface {
	HandleEvent(event Event)
}

type EventDispatcher struct {
	eventHandlers []EventHandler
}

func (d *EventDispatcher) RegisterHandler(handler EventHandler) {
	d.eventHandlers = append(d.eventHandlers, handler)
}

func (d *EventDispatcher) DispatchEvent(event Event) {
	for _, handler := range d.eventHandlers {
		handler.HandleEvent(event)
	}
}

type ExampleEventHandler struct {
	Name string
}

func (h *ExampleEventHandler) HandleEvent(event Event) {
	fmt.Printf("[%s] Received Event: ID=%d, Timestamps=%s\n", h.Name, event.ID, event.Timestamp.String())
}

func main() {
	dispatcher := EventDispatcher{}
	handler1 := &ExampleEventHandler{Name: "Handler 1"}
	handler2 := &ExampleEventHandler{Name: "Handler 2"}

	dispatcher.RegisterHandler(handler1)
	dispatcher.RegisterHandler(handler2)

	for i := 1; i < 10; i++ {
		event := Event{ID: i, Timestamp: time.Now()}
		dispatcher.DispatchEvent(event)
		time.Sleep(time.Second)
	}

}
