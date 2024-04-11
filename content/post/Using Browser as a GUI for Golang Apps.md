---
title: "Using Browser as a GUI for Golang Apps"
date: 2024-03-18
aliases: []
status:
id: 20240411164548351
---

# Using Browser as a GUI for Golang Apps

- Premise: Browser is a good enough GUI for most apps.



```js
// TODO: Make wsEndPt configurable
var wsEndPt = "ws://127.0.0.1:8000/ws; 


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
