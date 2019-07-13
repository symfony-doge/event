// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package event

import (
	"os"
)

// EventType is the category (or tag/marker) for events routing.
type EventType uint8

// Event represents a message in pubsub pattern with type information and
// a custom data payload.
type Event struct {
	// The event type to decide receiver's logic.
	Type EventType

	// A custom data payload.
	Payload interface{}
}

// Events is an alias to []Event
type Events []Event

// WithTypeAndPayload is a constructor that returns a new event with
// specified type and data payload.
func WithTypeAndPayload(t EventType, payload interface{}) Event {
	return Event{t, payload}
}

// MustListen is a shortcut for starting a common event listener session.
// It will abort the program execution if any error occurs.
func MustListen(fn ConsumeFunc) ROListenerSession {
	var listener = DefaultListenerInstance()

	listenerSession, listenErr := listener.Listen(fn)
	if nil != listenErr {
		fmt.Println("An error has been occurred during event.MustListen call:", listenErr)

		os.Exit(1)
	}

	return listenerSession
}
