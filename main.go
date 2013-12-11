//I am the entry point and interface for configuring PolySocket
package main

import (
  "flag"
  "strconv"
)

func main() {
  port := flag.Int("port", 8080, "port number for PolySocket daemon to listen on")

  flag.Parse()

  CreateServer(":" + strconv.Itoa(*port))
}
