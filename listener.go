// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package event

// ConsumeFunc is a callback signature for managing received events.
// It receives a read-only copy of an event from listening session.
type ConsumeFunc func(Event)

// Listener receives events (e.g. from workers) and calls a specified
// closure for processing. The Listen method returns a read-only listening
// session instance that exposes a communication channel between sender(s)
// and related event listener.
//
// Usage example:
//
// listenerSession, listenErr := eventListener.Listen(func(e Event) {})
// if listenErr != nil {
//     // Handle error.
// }
//
// var notifyChannel chan<- Event = listenerSession.NotifyChannel()
// notifyChannel <- Event{}
//
// listenerSession.Close()    // Use this method to close the channel safely.
type Listener interface {
	Listen(ConsumeFunc) (ROListenerSession, error)
}
