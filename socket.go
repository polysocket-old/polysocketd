package main

// socket contains the live websocket and functions to write to it and emit channel events out of the socket
import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"net/url"
)

func main() {
	target, err := url.Parse("ws://echo.websocket.org:80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(target.Host)
	header := http.Header{}
	conn, err := net.Dial("tcp", target.Host)
	if err != nil {
		log.Fatal(err)
	}
	ws, _, err := websocket.NewClient(conn, target, header, 512, 512)
	if err != nil {
		log.Fatal(err)
	}
	err = ws.WriteMessage(websocket.TextMessage, []byte("test"))
	if err != nil {
		log.Fatal(err)
	}
	messageType, message, err := ws.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("message type: %s\n", messageType)
	fmt.Println(string(message[:]))
}

type socket struct {
	id uuid.UUID
	ws *websocket.Conn
}

func NewSocket() {
}

// only string messages are considered right now

// here's how I'm thinking of "emitting"
// socket.C (a channel)
// - websocket has some data
// - socket.C <- websocket.data (which will block so this is in its own goroutine)
// FROM SERVER
// on /polysocket/jsonp?socket=#{this_socket}
// - select {
//      msg := socket.C
//      timeout := time.After
//   this lets us "block" hold open the jsonp connection until there is some data
