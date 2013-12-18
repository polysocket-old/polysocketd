package polysocket

import "fmt"

type Server struct {
	Addr string
}

func NewServer(port int) *Server {
	s := new(Server)
	s.Addr = fmt.Sprintf(":%v", port)

	return s
}
