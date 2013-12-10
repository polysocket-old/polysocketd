// I am the http server. I hold the routes, and interact with sockets
package main

import (
  "fmt"
  "github.com/drone/routes"
  . "net/http"
)

func main() {
  handler := routes.New()

  handler.Get("/", index)

  handler.Post("/polysocket/create", createSocket)

  handler.Post("/polysocket/send", sendMessage)

  handler.Get("/polysocket/:method(xhr|jsonp)", listenForMessages)

  Handle("/", handler)

  ListenAndServe(":8000", nil)
}

func index(w ResponseWriter, r *Request) {
  fmt.Fprintf(w, "We could put statistics here. That would be neat")
}

func createSocket(w ResponseWriter, r *Request) {
  fmt.Fprintf(w, "Don't mind me, just creatin a socket. Isn't that nice?")
}

func sendMessage(w ResponseWriter, r *Request) {
  fmt.Fprintf(w, "I'm gonna send that message for you")
}

func listenForMessages(w ResponseWriter, r *Request) {
  method := r.URL.Query().Get(":method")

  fmt.Fprintf(w, "Gonna listen for messages using: %s", method)
}
