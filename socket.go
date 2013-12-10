// socket contains the live websocket and functions to write to it and emit channel events out of the socket

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
