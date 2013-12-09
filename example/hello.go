package main

import (
	polysocket "github.com/polysocket/polysocket-relay"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	_ = polysocket.PolySocket{}
	http.HandleFunc("/long", func(w http.ResponseWriter, req *http.Request) {
		go func() {
			timeout, err := polysocket.Timeout(req, "timeout")
			if err != nil {
				io.WriteString(w, "bad timeout!")
				return
			}
			log.Print(timeout)
			select {
			case <-time.After(timeout):
				io.WriteString(w, "hellur")
			}
		}()
	})
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
