# polysocket-relay

I transform xhr-streaming and jsonp-polling into beautiful raw websocket connections.

## api

### `POST /polysocket/create`

This creates an open socket on the polysocket server for you to talk through. If the response is not 200, then there was an issue with the server, with networking, with load, or with you not giving enough parameters.

If the response is 200 but not ok, then we tried to establish a websocket connection for you to your target server, but that server had a websocket error. This should be treated like an error in creating a `new WebSocket` so a server that rejects or errors the connection.

The response gives you a `psid` (polysocket id) which is your session number for continuing to speak and receive on this socket.

The response also gives you a `relay` which is a hostname that you should continue talking to. This lets one server handle your rest calls and maintain state (e.g. 24.relay.polysocket.io).

**params**

* `target` (String) valid WebSocket URL, who you're connecting to through the relay

**response**
`400` bad request means we couldn't attemp to fulfil the request because you didn't provide necessary parameters
`200` means server has handled your request without issue, but your response may still be an error

* `ok` (Boolean) true when no error
* `error` (String, optional) present when ok is false
* `psid` (String) your polysocket id
* `relay` (String) the hostname of the relay you should talk to

## LICENSE

MIT

