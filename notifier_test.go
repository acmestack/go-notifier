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
	"testing"
	"time"
)

func TestNotifier(t *testing.T) {
	type fields struct {
		observers map[string][]*Event
	}
	type args struct {
		name string
		e    *Event
	}
	
	tests := []struct {
		name   string
		fields fields
		args   []args
	}{
		{
			name: "first",
			args: []args{
				{
					name: "sense1",
					e: &Event{
						Name: "a",
						Watch: func(e *Event, data any) {
							t.Logf("data %v", data)
						},
					},
				},
				{
					name: "sense2",
					e: &Event{
						Name: "b",
						Watch: func(e *Event, data any) {
							t.Logf("data %v", data)
						},
					},
				},
			},
		},
		{
			name: "second",
			args: []args{
				{
					name: "sense1",
					e: &Event{
						Name: "a",
						Watch: func(e *Event, data any) {
							t.Logf("data %v", data)
						},
					},
				},
				{
					name: "sense2",
					e: &Event{
						Name: "b",
						Watch: func(e *Event, data any) {
							t.Logf("data %v", data)
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &observer{
				observers: tt.fields.observers,
			}
			for i, arg := range tt.args {
				s.register(tt.name, arg.e)
				if s.observers[tt.name][i] != arg.e {
					t.Fatal("not eq")
				}
			}
			
			s.notification(tt.name, time.Now())
			
			for _, arg := range tt.args {
				s.unregister(tt.name, arg.e)
			}
		})
	}
}
