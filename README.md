# Symfony Doge's Event

A set of reusable components for building a simple message-oriented middleware in Go.

## Installation

```
$ go get -u -d github.com/symfony-doge/event@latest
```

## Usage

### DefaultListener

[DefaultListener](default_listener.go) acts like a subscriber that receives and process events (i.e. messages in context of pubsub pattern)
from multiple publishers. It listens a channel wrapped by [ROListenerSession](ro_listener_session.go).
This implementation doesn't support any custom routing.

See [example](example/one_subscriber_many_publishers.go) code snippet:

```go
fmt.Println("One subscriber many publishers example...")

var listener event.Listener = event.DefaultListenerInstance()

fmt.Println("Starting listening session...")

var consumeFunc event.ConsumeFunc = func (e event.Event) {
	fmt.Printf("An event has received. Type: %d, Payload: %v\n", e.Type, e.Payload)
}

listenerSession, listenErr := listener.Listen(consumeFunc)
if nil != listenErr {
	fmt.Errorf("An error has occurred during Listen call:", listenErr)
}

fmt.Println("Scheduling cleanup for listening session...")

defer listenerSession.Close()

var notifyChannel chan<- event.Event = listenerSession.NotifyChannel()

fmt.Println("Pushing events to the notification channel...")

notifyChannel <- event.WithTypeAndPayload(1, "test payload 1")
notifyChannel <- event.WithTypeAndPayload(2, "test payload 2")
notifyChannel <- event.WithTypeAndPayload(3, "test payload 3")
```

Output will be:

```
One subscriber many publishers (no routing) example...
Starting listening session...
Scheduling cleanup for listening session...
Pushing events to the notification channel...
Event has received. Type: 1, Payload: test payload 1
Event has received. Type: 2, Payload: test payload 2
Event has received. Type: 3, Payload: test payload 3
```

## See also

- [agoalofalife/event](https://github.com/agoalofalife/event) — The Observer pattern implementation in Go.
- [olebedev/emitter](https://github.com/olebedev/emitter) — Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins.
- [leandro-lugaresi/hub](https://github.com/leandro-lugaresi/hub) — A fast Message/Event Hub using publish/subscribe pattern with support for topics like* rabbitMQ exchanges for Go applications.
- [asaskevich/EventBus](https://github.com/asaskevich/EventBus) — Lightweight eventbus with async compatibility for Go.

## Changelog

All notable changes to this project will be documented in [CHANGELOG.md](CHANGELOG.md).
