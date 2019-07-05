// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package event

// ROListenerSession is a read-only listener session that's used
// as a bridge between event listener and sender(s) (e.g. worker); a caller
// should pass the notification channel associated with listening session
// to the sender(s) to start communication via events.
type ROListenerSession struct {
	// The base listener session implementation.
	base *listenerSession

	// A channel that should be used by senders to push their events.
	// If a notification channel becomes closed, listening session ends.
	// Notification channel must be closed only by Close method, it is not
	// safe to use a built-in close function directly due to non-blocking
	// event processing.
	notifyChannel chan<- Event
}

// NotifyChannel method returns a notification channel for sending events.
func (ls ROListenerSession) NotifyChannel() chan<- Event {
	return ls.notifyChannel
}

// Close method should be used to safely close the notification channel and
// stop the event listening session.
func (ls ROListenerSession) Close() {
	// This will ensure that all remaining events are properly processed.
	defer ls.wait()

	close(ls.notifyChannel)
}

func (ls ROListenerSession) wait() {
	<-ls.base.done
}

func newROListenerSession(base *listenerSession, nc chan<- Event) ROListenerSession {
	return ROListenerSession{base, nc}
}
