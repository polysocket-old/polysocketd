package main

import (
  "code.google.com/p/go-uuid/uuid"
  "fmt"
  "github.com/gorilla/websocket"
  "log"
  "net"
  "net/http"
  "net/url"
  "time"
)

func main() {
  target, err := url.Parse("ws://echo.websocket.org:80")
  if err != nil {
    log.Fatal(err)
  }
  header := http.Header{}
  config := socketConfig{
    target:  target,
    header:  header,
    timeout: time.Second * 5,
  }
  s, _ := NewSocket(config)
  go func() {
    s.WriteMessage <- "lulz"
  }()
  <-s.ReadMessage

}

type socket struct {
  id           string
  ws           *websocket.Conn
  ReadMessage  chan string
  WriteMessage chan string
}

type socketConfig struct {
  header  http.Header
  target  *url.URL
  timeout time.Duration
}

func NewSocket(config socketConfig) (*socket, error) {
  conn, err := net.Dial("tcp", config.target.Host)
  if err != nil {
    return nil, err
  }
  ws, _, err := websocket.NewClient(conn, config.target, config.header, 512, 512)
  if err != nil {
    return nil, err
  }
  s := &socket{uuid.New(), ws, make(chan string), make(chan string)}
  go func() {
    for {
      message := <-s.WriteMessage
      ws.WriteMessage(websocket.TextMessage, []byte(message))
    }
  }()
  go func() {
    for {
      messageType, message, err := ws.ReadMessage()
      if err == nil {
        fmt.Printf("%s\n%s", messageType, message)
        s.ReadMessage <- string(message[:])
      }
    }
  }()
  return s, nil
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
