// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package event

import (
	"sync"
)

const (
	// Max session contexts for default listener's history field.
	dlMaxHistoryEntries = 10

	// Buffer size for events channel.
	// In ideal case we should not block senders during their communication
	// routines with listener, but it depends on concrete task and how often
	// they will communicate.
	dlNotifyChannelBufferSize = 1 << 3
)

var defaultListenerInstance *DefaultListener

var defaultListenerOnce sync.Once

// DefaultListener encapsulates a common event listening logic (one subscriber
// to many publishers; a single channel will be used as a communication bridge)
// and supports multiple listening sessions.
type DefaultListener struct {
	// Listening history.
	history []*listenerSession
}

// Starts and returns new listening session with notification channel to which
// senders should push their events; stops listening when the notify channel
// becomes closed.
func (l *DefaultListener) Listen(fn ConsumeFunc) (ROListenerSession, error) {
	var notifyChannel = make(chan Event, dlNotifyChannelBufferSize)
	var listenerSession = newListenerSession()

	defer l.historyRotate(listenerSession)

	go func() {
		for {
			select {
			// Blocks default flow whenever a new event is available to consume
			// or channel becomes closed.
			case event, isChannelOpen := <-notifyChannel:
				if !isChannelOpen {
					// We should set channel to nil, to ensure it will not
					// block default case with endless communication loop
					// (closed channels fires immediately).
					notifyChannel = nil

					break
				}

				listenerSession.received = append(listenerSession.received, event)
			// While waiting for new events, we use goroutine time
			// to process ones which already received (non-blocking approach).
			default:
				// We still have to process events until each will be
				// "consumed", then we can change session state, but notify
				// channel is safe to be closed earlier.
				if listenerSession.consumed < len(listenerSession.received) {
					var next Event = listenerSession.received[listenerSession.consumed]

					fn(next)
					listenerSession.consumed++

					break
				}

				if nil == notifyChannel {
					listenerSession.done <- true

					return
				}
			}
		}
	}()

	var readOnlyListenerSession = newROListenerSession(listenerSession, notifyChannel)

	return readOnlyListenerSession, nil
}

func (l *DefaultListener) historyRotate(ls *listenerSession) {
	if len(l.history) >= dlMaxHistoryEntries {
		l.history = make([]*listenerSession, dlMaxHistoryEntries)
	}

	l.history = append(l.history, ls)
}

func NewDefaultListener() *DefaultListener {
	return &DefaultListener{
		history: make([]*listenerSession, dlMaxHistoryEntries),
	}
}

func DefaultListenerInstance() *DefaultListener {
	defaultListenerOnce.Do(func() {
		defaultListenerInstance = NewDefaultListener()
	})

	return defaultListenerInstance
}
