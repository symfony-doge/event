// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package event

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
