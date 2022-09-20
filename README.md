# Hello conspire

## Running

Running the current state.

### Start WS Server

```shell
go run main/conspire.go servews
```

### Install a Static HTTP File Server

(If you haven't already.)

```shell
npm install -g live-server
```

### Start HTTP Server

The 'client' is script on the `index.html` page that the server serves up.

```shell
cd http/wsserver
live-server
```

## TODO

* ws: A handler for WinTest
* ws: A generic wss:// handler for testing clients
* ws: More messages back and forth
* ws: Handle multiple clients

### Features TODO

* A general-purpose wss:// handler to develop client ws apps against
  * What data?

