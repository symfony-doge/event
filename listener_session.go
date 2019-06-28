// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package event

// listenerSession is not designed to be a thread-safe unit
// and should not be accessed from multiple execution flows, i.e.
// only listener and internals can access and manage it.
type listenerSession struct {
	// Holds all received events within a single listening session.
	received Events

	// Index of the next event to be "consumed", i.e. processed with ConsumeFunc.
	consumed int

	// Receives a value when listening is done.
	done chan bool
}

func newListenerSession() *listenerSession {
	return &listenerSession{
		received: Events{},
		consumed: 0,
		done:     make(chan bool, 1),
	}
}
