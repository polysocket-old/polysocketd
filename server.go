// I am the http server. I hold the routes, and interact with sockets
package main

import (
  "fmt"
  "github.com/gorilla/mux"
  . "net/http"
)

type PolysocketServer struct {
  port        string
  isListening bool
}

func (s PolysocketServer) CreateServer(port string) PolysocketServer {
  if s.isListening {
    return s
  }

  handler := mux.NewRouter()

  handler.HandleFunc("/", index).Methods("GET")

  handler.HandleFunc("/polysocket/polysocket.js", sendClientLibrary).Methods("GET")

  handler.HandleFunc("/polysocket/create", createSocket).Methods("POST")

  handler.HandleFunc("/polysocket/send", sendMessage).Methods("POST")

  handler.HandleFunc("/polysocket/{method:(xhr|jsonp)}", listenForMessages).Methods("GET")

  Handle("/", handler)

  ListenAndServe(s.port, nil)

  s.isListening = true

  return s
}

func index(w ResponseWriter, r *Request) {
  fmt.Fprintf(w, "We could put statistics here. That would be neat")
}

func sendClientLibrary(w ResponseWriter, r *Request) {
  fmt.Fprintf(w, "One day I'll give you a javascript file for your browser.")
}

func createSocket(w ResponseWriter, r *Request) {
  fmt.Fprintf(w, "Don't mind me, just creatin a socket. Isn't that nice?")
}

func sendMessage(w ResponseWriter, r *Request) {
  fmt.Fprintf(w, "I'm gonna send that message for you")
}

func listenForMessages(w ResponseWriter, r *Request) {
  method := mux.Vars(r)["method"]

  fmt.Println(mux.Vars(r))

  fmt.Fprintf(w, "Gonna listen for messages using: %s", method)
}
