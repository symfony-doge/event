// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package example

import (
	"fmt"
	"os"

	"github.com/symfony-doge/event"
)

func osmpConsumeFunc(e event.Event) {
	fmt.Printf("An event has been received. Type: %d, Payload: %v\n", e.Type, e.Payload)
}

func OneSubscriberManyPublishers() {
	fmt.Println("One subscriber many publishers (no routing) example...")
	var listener event.Listener = event.DefaultListenerInstance()

	fmt.Println("Starting listening session...")
	listenerSession, listenErr := listener.Listen(osmpConsumeFunc)
	if nil != listenErr {
		fmt.Println("An error has been occurred during Listen call:", listenErr)
		os.Exit(1)
	}

	fmt.Println("Scheduling cleanup for listening session...")
	defer listenerSession.Close()

	var notifyChannel chan<- event.Event = listenerSession.NotifyChannel()

	fmt.Println("Pushing events to the notification channel...")
	notifyChannel <- event.WithTypeAndPayload(1, "test payload 1")
	notifyChannel <- event.WithTypeAndPayload(2, "test payload 2")
	notifyChannel <- event.WithTypeAndPayload(3, "test payload 3")
}
