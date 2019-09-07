package main

import (
	"fmt"
	"time"
)

type (
	Event struct {
		Data int64
	}

	Observer interface {
		OnNotify()
	}

	Notifier interface {
		Register(observer Observer)
		Notify()
	}
)

type (
	eventObserver1 struct {
		event *Event
		id    int
	}

	eventObserver2 struct {
		event *Event
		id    int
	}

	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (n *eventNotifier) Register(observer Observer) {
	n.observers[observer] = struct{}{}
}

func (n *eventNotifier) Notify() {
	for o := range n.observers {
		o.OnNotify()
	}
}

func (o *eventObserver2) OnNotify() {
	fmt.Printf("Observer %d recieved: %d\n", o.id, o.event.Data)
}

func (o *eventObserver1) OnNotify() {
	fmt.Printf("Observer %d recieved: %d\n", o.id, o.event.Data)
}

func main() {
	notifier := eventNotifier{
		observers: map[Observer]struct{}{},
	}
	event := &Event{0}

	notifier.Register(&eventObserver1{
		event: event,
		id:    1,
	})

	notifier.Register(&eventObserver2{
		event: event,
		id:    2,
	})

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C

	for {
		select {
		case <-stop:
			return

		case t := <-tick:
			event.Data = t.UnixNano()
			notifier.Notify()
		}

	}
}
