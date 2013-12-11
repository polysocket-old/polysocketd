// I am the http server. I hold the routes, and interact with sockets
package main

import (
  "fmt"
  "github.com/gorilla/mux"
  . "net/http"
)

type PolySocketServer struct{}

func (s PolySocketServer) CreateServer(port string) PolySocketServer {
  handler := mux.NewRouter()

  handler.HandleFunc("/", index).Methods("GET")

  handler.HandleFunc("/polysocket/polysocket.js", sendClientLibrary).Methods("GET")

  handler.HandleFunc("/polysocket/create", createSocket).Methods("POST")

  handler.HandleFunc("/polysocket/send", sendMessage).Methods("POST")

  handler.HandleFunc("/polysocket/{method:(xhr|jsonp)}", listenForMessages).Methods("GET")

  Handle("/", handler)

  ListenAndServe(port, nil)

  return s
}

func index(w ResponseWriter, r *Request) {
  fmt.Fprintf(w, "We could put statistics here. That would be neat")
}

func sendClientLibrary(w ResponseWriter, r *Request) {
  fmt.Fprintf(w, "One day I'll give you a javascript file for your browser.")
}

func createSocket(w ResponseWriter, r *Request) {
  params := []string{"target", "origin"}

  if !validateQueryString(w, r, params) {
    return
  }

  fmt.Fprintf(w, "Don't mind me, just creatin a socket. Isn't that nice?")
}

func sendMessage(w ResponseWriter, r *Request) {
  params := []string{"socket", "events"}

  if !validateQueryString(w, r, params) {
    return
  }

  fmt.Fprintf(w, "I'm gonna send that message for you")
}

func listenForMessages(w ResponseWriter, r *Request) {
  method := mux.Vars(r)["method"]

  fmt.Fprintf(w, "Gonna listen for messages using: %s", method)
}

func validateQueryString(w ResponseWriter, r *Request, params []string) bool {
  queryParams := r.URL.Query()

  for _, param := range params {

    if queryParams[param] == nil {
      errString := "Missing one or more parameters: "

      for _, param := range params {
        errString += (param + " ")
      }

      Error(w, errString, 400)
      return false
    }

  }
  return true
}
