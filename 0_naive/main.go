// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "fmt"

const str = "hi there!"

type Message string

func NewMessage() Message {
	return Message(str + " :called by NewMessage()")
}

func NewGreeter(m Message) Greeter {
	return Greeter{
		Msg: m,
	}
}

type Greeter struct {
	Msg Message // <- adding a Message field
}

func (g Greeter) Greet() Message {
	return g.Msg + " :called by g.Greet()"
}

func NewEvent(g Greeter) Event {
	return Event{
		Grer: g,
	}
}

type Event struct {
	Grer Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Grer.Greet()
	fmt.Println(msg + " :called by e.Start()")
}

func main() {
	m := NewMessage()
	g := NewGreeter(m)
	e := NewEvent(g)

	e.Start()
}
