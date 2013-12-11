//I am the entry point and interface for configuring PolySocket
package main

import (
  "flag"
  "strconv"
)

func main() {
  port := flag.Int("port", 8080, "port number for PolySocket daemon to listen on")

  flag.Parse()

  s := PolySocketServer{}
  s.CreateServer(":" + strconv.Itoa(*port))
}
