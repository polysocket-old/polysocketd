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

### `send`  
### `JSONP {uniq}.tom.polysocket.com/s/#{socket}`  
### `POST {uniq}.tom.polysocket.com/s/#{socket}`  

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

### `JSONP {uniq}.tom.polysocket.com/j/#{socket}`  

this starts a jsonp long-polling call. this is how you receive data out of your socket.

**params**

```javascript
{
  cb  : (string)
  ttl : (integer) time before request expires and server should return with no data (in milliseconds)
}
```

**response**

`400` bad request means bad parameters

`403` unauthorized means your socket id could not be found

`200` means the server has a valid response for you to process

```javascript
{
  ok       : true,
  messages : [
    {
      type: 'heartbeat' // heartbeat
    }
    {
      type: 'text'
      data: (string)
    }
    {
      type: 'binary'
      data: (base64 string)
    }
  ]
}
{
  error : (string)
  ok    : false
}
```

## dns

each polysocketd server should have a valid fqdn, e.g. tom.polysocket.com. additionally, each fqdn should handle all subdomains, e.g. 1234.tom.polysocket.com.

this allows the polysocket relay servers to tell a client to connect to a specific polysocketd server (the one that has opened a websocket and ready to send/receive). this also allows the browser to open a second polysocket connection to the same server without holding too many browser connections to the same domain (connecting once to 1234.tom.polysocket.com and once to 4321.tom.polysocket.com).

router servers can be round-robined and made highly available. they are only used in negotiating a new socket connection and doing authentication, e.g. router.polysocket.com should dns rr to two unique addresses for availability.

## license

MIT

