---
title: 
date: 2024-04-11T16:45:48+05:30
lastmod: 2024-04-11T20:00:18+05:30
author: Riz

description: 
categories: []
tags: []

draft: false
enableDisqus : false
enableMathJax: false
disableToC: false
disableAutoCollapse: true
---



<div id="notice">NOTE: Error handling in code blocks removed for brevity sake.</div>

- Premise: Browser is a good enough GUI for most apps.

So this is how I set out to overcome some basic annoyances in using your installed browser as a reasonable front-end to golang apps. The annoyances being
- When I double-click on executable or start it from command line, it should start the backend server and open the localhost address automatically using your default browser or configured option.
- When I close the relevant tab/tabs or browser, the backend server should shutdown automatically.

So first step is rather easy - just tell the server to open the address after it starts listening and before it starts serving thereby blocking the thread. That means instead of the usual 

```go
	srv.ListenAndServe()
```
you will write it as
```go
listener, _ := net.Listen("tcp", srv.Addr) 

// import "github.com/skratchdot/open-golang/open"
_ = open.Start("http://" + srv.Addr) 

srv.Serve(listener)
```
This starts the listener, opens up the localhost address in default browser and then start serving. So far I did not find the need to introduce a sleep delay between the opening the address and starting ther server, but that can be considered if your browser start-up times are considerable. 

Now the possible variation here is usage of `os.exec` in place of `open.Start` to control the browser itself. Firefox offers command-line options like

```shell
firefox-developer-edition --new-window --kiosk "http://localhost:1313/"
```
which will open a new window any of the usual browser window decorations. You can quit using `Ctrl+q`.  If that is too extreme, there are [add-ons](https://addons.mozilla.org/en-GB/firefox/addon/popup-window/), [configurations](https://support.mozilla.org/en-US/questions/1292666) and userChrome CSS tricks which will offer you similar abilities. Of course you will limit these modifications to a particular profile and choose that profile from commandline. I am sure chrome offers similar capabilities.

Now the second part is trickier. So here is what we need to achieve
1. Server should start serving pages.
2. As long as there are pages served by our server being accessed - server should stay running.
3. When all the pages are closed, server should shutdown.

The problem is, we need the pages to inform that they are being accessed. One way to do this is to send a health check signal to server from pages with a timer. But if we are injecting custom javascript anyway, might as well do it properly. I am talking about websockets. 

So here is what we are going to do.
1. Block the shutting down of server using a channel
2. On the client side, inform the server when they open as well as close via a dedicated endpoint.
4. Each new page increments the number of connections by 1
5. Each page closure decrements the number of connections by 1 and checks for existing number of connections. If the number of connections is zero, send signal to shutdown server via channel. 

So step 1: Block the shutting down
```go
// Declare a global variable
var endofSession = make(chan bool)

```
```go
// Start the shutdown in a separate goroutine
go func() {
	<-endofSession
	srv.Shutdown(context.Background())
}()
```



Step 2: Let the client inform server upon opening a new connection as well as closing
 For that, the following js should to inserted to all relevant pages. On the server side, add CORS to ensure no one else is pinging the endpoint.
```js
// TODO: Make wsEndPt configurable
var wsEndPt = "ws://127.0.0.1:1313/ws; 


let socket = new WebSocket(wsEndPt);

socket.onopen = () => {
  console.log("Registering Client");
  socket.send("Connection from Client ID: ${browser.tabs.Tab.id}");
};

socket.onclose = (event) => {
  console.log("Socket Closed Connection: ", event);
  socket.send("Client Closed!");
};
// TODO: On message here
socket.onerror = (error) => {
  console.log("Socket Error: ", error);
};
```

Now step 3 and 4 on the server side

Set aside the endpoint
```go
rtr.HandleFunc("/ws", handleWs)
```

Handle the ping.
```go
//import github.com/gorilla/websocket"
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var currConn = 0

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func handleWs(w http.ResponseWriter, r *http.Request) {
	currConn = currConn + 1
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // TODO: Check if it is working
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("New Connection"))
	if err != nil {
		log.Println(err)
	}
	reader(ws) // Block the flow till client disconnects

	// Once client disconnects reduce the number of connections by one
	currConn = currConn - 1
	if currConn > 0 {
		return
	}
	// Wait 10secs for any more connection and any operations to finish
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		if currConn != 0 {
			return
		}
	}
	endofSession <- true
}
```