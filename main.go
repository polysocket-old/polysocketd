package main

import "github.com/polysocket/polysocketd/polysocket"

func main() {
	_ = polysocket.NewServer(8080)
}
