# Symfony Doge's Event

A set of reusable components for building a simple message-oriented middleware in Go.

## Installation

```
$ go get -u -d github.com/symfony-doge/event@latest
```

## Usage

### DefaultListener

[DefaultListener](default_listener.go) acts like a subscriber that receives and process events (i.e. messages in context of pubsub pattern)
from multiple publishers. It listens a channel wrapped by [listenerSession](listener_session.go).

## See also

- [agoalofalife/event](https://github.com/agoalofalife/event) — The Observer pattern implementation in Go.
- [olebedev/emitter](https://github.com/olebedev/emitter) — Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins.
- [leandro-lugaresi/hub](https://github.com/leandro-lugaresi/hub) — A fast Message/Event Hub using publish/subscribe pattern with support for topics like* rabbitMQ exchanges for Go applications.
- [asaskevich/EventBus](https://github.com/asaskevich/EventBus) — Lightweight eventbus with async compatibility for Go.

## Changelog

All notable changes to this project will be documented in [CHANGELOG.md](CHANGELOG.md).
