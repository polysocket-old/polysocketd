# polysocket

just use websockets!

most browsers support websockets, and the ones that don't polysocket will upgrade them! so, STAP using non-standard protocols for implementing realtime! let's start building libraries and utilities on top of the standard for realtime: websockets.

## api

### `JSONP router.polysocket.com/create`

jsonp since we're aiming for supporting all browsers. old ones only have jsonp as an option (since it's a cross-origin request), and new ones may use jsonp.

```
// pseudocode

// validate request parameters
ensure target is websocket url
ensure origin is valid with key
ensure caller is present

// validate authentication
token authorizes target websocket url
within connection limit for token

// which polysocket server should handle this client?
get polysocketd fqdn that is most available

// establish socket on that remote server
socket, err = polysocket.establish(target) // redis message passing
```

**params**

```javascript
{
  cb     : (string) jsonp callback (also serves as cache-buster)
  from   : (string) e.g. 'polysocketjs-v0.0.1'
  target : (string : valid websocket uri)
  token  : (string) auth token
}
```

**response**

```javascript
{
  ok          : true
  polysocketd : '{uniq}.tom.polysocket.com'
  socket      : (string) this is your socket and your session id for communication
}
{
  error : (string)
  ok    : false
}
```

### `JSONP {uniq}.tom.polysocket.com/send/#{socket}`
### `POST {uniq}.tom.polysocket.com/send/#{socket}`

this is how you write data to your socket.

**params**

```javascript
{
  cb : (string) only for jsonp version
  d  : (string) base64 string when t = 'b', or utf8 string
  t  : (string : ['b','t']) binary, or text // TODO handle continuation for when payload longer than GET URL capacity
}
```

**reponse**

`400` bad request means you are missing parameters or they are poorly formatted

`403` unauthorized means your socket isn't valid

`201` means we have accepted your data and pushed it along your socket, feel free to send more now

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

