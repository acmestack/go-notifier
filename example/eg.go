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

package main

import (
	"fmt"
	"sync"
	"time"
	
	"github.com/acmestack/go-notifier"
	_ "github.com/acmestack/go-notifier/example/another"
	"github.com/acmestack/go-notifier/example/constant"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	notifier.Register(constant.App, &notifier.Event{
		Name: fmt.Sprintf("event name: %d", id),
		Watch: func(event *notifier.Event, data any) {
			fmt.Printf("event name %s data %v\n", event.Name, data)
		},
	})
}

func main() {
	var wg sync.WaitGroup
	
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	
	wg.Wait()
	
	notifier.Notification(constant.App, time.Now())
}
