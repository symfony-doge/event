// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package event

type EventType uint8

// Represents an event with type information and a custom payload.
type Event struct {
	// The event type to decide receiver's logic.
	Type EventType

	// A custom data payload.
	Payload interface{}
}

type Events []Event

func NewEvent(t EventType, payload interface{}) Event {
	return Event{t, payload}
}
