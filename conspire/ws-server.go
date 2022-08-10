package conspire

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/gorilla/websocket"
  "log"
  "net/http"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
}

func WsMain() {
  setupRoutes()

  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal(err)
  }
}

func setupRoutes() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/ws", wsEndpoint)
}

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
  upgrader.CheckOrigin = func(r *http.Request) bool { return true }

  ws, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Client Successfully Connected...")

  reader(ws)
}

func reader(conn *websocket.Conn) {
  for {
    messageType, p, err := conn.ReadMessage()
    if err != nil {
      fmt.Println(err)
      return
    }

    fmt.Println(string(p))
    if err := conn.WriteMessage(messageType, p); err != nil {
      fmt.Println(err)
      return
    }
  }
}


