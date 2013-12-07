# polysocket-relay

I transform xhr-streaming and jsonp-polling into beautiful raw websocket connections.

## api

### `POST /polysocket/create`

This creates an open socket on the polysocket server for you to talk through. If the response is not 200, then there was an issue with the server, with networking, with load, or with you not giving enough parameters.

If the response is 200 but not ok, then we tried to establish a websocket connection for you to your target server, but that server had a websocket error. This should be treated like an error in creating a `new WebSocket` so a server that rejects or errors the connection.

The response gives you a `socket` which is your session number for continuing to speak and receive on this socket.

The response also gives you a `relay` which is a hostname that you should continue talking to. This lets one server handle your rest calls and maintain state (e.g. 24.relay.polysocket.io).

**params**

```javascript
{
  target : (String) valid WebSocket URL, who you're connecting to through the relay
}
```

**response**

`400` bad request means we couldn't attemp to fulfil the request because you didn't provide necessary parameters

`200` means server has handled your request without issue, but your response may still be an error

```javascript
{
  ok     : (Boolean) true when no error
  error  : (String, optional) present when not ok
  socket : (String) your socket id
  relay  : (String) hostname of the relay you should talk to
}
```

### `POST ://#{relay}/polysocket/send`

This pushes data into you established socket. This is a call independent of your jsonp long polling and xhr-streaming. This is your way to send data, the other calls are how you receive data. Combined to make full-duplex!

**params**

```javascript
{
  socket : (String) your established socket id
  events : [
    {
      type : 'string' // string type of message send
      data : '1234'
    },
    {
      type : 'binary'
      data : 'ABCA' // base64 binary
    }
  ]
}
```

**reponse**

`400` bad request means you are missing parameters or they are poorly formatted

`403` unauthorized means your socket isn't valid

`201` means we have accepted your data and pushed it along your socket

There is no body to the response. Just a 201 code. Once you receive this code, you are free to POST again. You shouldn't POST multiple times before receiving a response to ensure that order is maintained.

### `GET ://#{relay}/polysocket/jsonp?socket=#{socket}&timeout=#{timeout_ms}&callback=#{callback_fn}`

This starts a jsonp long-polling call for receiving data over your socket. This will timeout and return with no data after timeout elapses and no data was sent. This gives the browser control over the timeout. It should be set to a time less than the browser deeming the connection as "timed out" (less than 30 seconds).

After you receive this response, you are expected to make a new call so you can get the next bit of data coming your way.

**response**

`400` bad request means bad parameters

`403` unauthorized means your socket id could not be found

`200` means the server has a valid response for you to process

```javascript
{
  ok     : (Boolean) true when no error
  error  : (String, optional) present when ok is false
  events : [ (array, zero or more in-order events)
    {
      type : 'heartbeat' (just the server telling you you're still alive, happens after a timeout)
    },
    {
      type : 'string' (you have a string message to process)
      data : '1234'
    },
    {
      type : 'binary' (you have a binary message to process - unsupported for now)
      data : 'ABC'
    }
  ]
}
```

## DNS

Each relay should have a valid FQDN. There should also be a round-robin (or other distribution method) endpoint which will route requests to an arbitrary relay server. Only the `/polysocket/create` method should hit the round-robin endpoint. All other requests should target a specific FQDN of a relay server.

## LICENSE

MIT

