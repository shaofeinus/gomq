package pubsub

import (
	"errors"
	"fmt"
	"github.com/shaofeinus/gomq"
)

var EVENTS = make(map[string]Event)

type Event struct {
	Name        string
	Exchange    string
	Subscribers map[string]Subscriber
}

type Subscriber struct {
	Name  string
	Queue string
	Fn    func(args map[string]interface{})
}

func RegisterEvent(name string) *Event {
	if _, ok := EVENTS[name]; !ok {
		EVENTS[name] = Event{
			Name:        name,
			Exchange:    fmt.Sprintf("pubsub-%s", name),
			Subscribers: map[string]Subscriber{},
		}
	}
	event := EVENTS[name]
	return &event
}

func Publish(event string, args map[string]interface{}) error {
	ev, err := findEvent(event)
	if err != nil {
		return err
	}
	return gomq.SendJSONToExchange(ev.Exchange, map[string]interface{}{
		"name": ev.Name,
		"args": args,
	})
}

func Subscribe(event *Event, subscriber string, fn func(args map[string]interface{})) {
	queue := fmt.Sprintf("pubsub-%s-%s", event.Name, subscriber)
	event.Subscribers[subscriber] = Subscriber{
		Name:  subscriber,
		Queue: queue,
		Fn:    fn,
	}
}

func WorkOnSubscriber(event string, subscriber string) error {
	ev, sub, err := findSubscriber(event, subscriber)
	if err != nil {
		return err
	}
	err = gomq.BindQueue(sub.Queue, ev.Exchange)
	if err != nil {
		return err
	}
	gomq.ConsumeJSON(sub.Queue, makeHandleEventJson(sub.Fn))
	return nil
}

func findEvent(name string) (*Event, error) {
	if event, ok := EVENTS[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Event not registered \"%s\"", name))
	} else {
		return &event, nil
	}
}

func findSubscriber(event string, subscriber string) (*Event, *Subscriber, error) {
	ev, err := findEvent(event)
	if err != nil {
		return nil, nil, err
	}
	if sub, ok := ev.Subscribers[subscriber]; !ok {
		return nil, nil, errors.New(fmt.Sprintf("Subscriber not registered \"%s\"", subscriber))
	} else {
		return ev, &sub, nil
	}
}

func makeHandleEventJson(fn func(args map[string]interface{})) func(eventJson map[string]interface{}) {
	return func(eventJson map[string]interface{}) {
		args := eventJson["args"].(map[string]interface{})
		fn(args)
	}
}
