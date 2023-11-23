/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package notifier

import (
	"sync"
)

type observer struct {
	lock      sync.RWMutex
	observers map[string][]*Event
}

var o *observer

func init() {
	o = new(observer)
}

// Register one event to notifier.
func Register(name string, e *Event) {
	o.register(name, e)
}

// UnRegister one event from notifier.
func UnRegister(name string, e *Event) () {
	o.unregister(name, e)
}

// Notification a event by name and data.
func Notification(name string, data any) {
	o.notification(name, data)
}

func (o *observer) register(name string, e *Event) {
	o.lock.Lock()
	defer o.lock.Unlock()
	
	if o.observers == nil {
		o.observers = make(map[string][]*Event, 10)
	}
	
	// override exist event by name
	for i, events := range o.observers[name] {
		if events.Name == e.Name {
			o.observers[name][i] = e
			return
		}
	}
	
	// append new event
	o.observers[name] = append(o.observers[name], e)
}

func (o *observer) unregister(name string, e *Event) {
	o.lock.Lock()
	defer o.lock.Unlock()
	
	for i, obs := range o.observers[name] {
		if obs == e {
			o.observers[name] = append(o.observers[name][:i], o.observers[name][i+1:]...)
			break
		}
	}
}

func (o *observer) notification(name string, data any) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	
	for _, events := range o.observers[name] {
		events.Watch(events, data)
	}
}
