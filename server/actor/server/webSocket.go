// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import "github.com/EtixLabs/cameradar/server/adaptor"

// WebSocket manages server communication using a websocket adaptor
type WebSocket struct {
	wsf adaptor.WebSocketFactory

	client   adaptor.WebSocket
	messages chan string
}

// NewWebSocket creates a Server actor that uses a WebSocketFactory
func NewWebSocket(
	wsf adaptor.WebSocketFactory,
) *WebSocket {
	wsServer := &WebSocket{
		wsf: wsf,

		messages: make(chan string),
	}
	go wsServer.Run()
	return wsServer
}

// Run starts to listen on messages
func (ws *WebSocket) Run() {
	for {
		msg := <-ws.messages
		ws.client.Write() <- msg
	}
}
