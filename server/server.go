package server

import (
	"log"
	"net/http"
)

type WebSocketServer struct {
	hub *Hub
	addr string
}

func NewWebSocketServer() *WebSocketServer{
	return &WebSocketServer{
		hub: newHub(),
	}
}


func (ws *WebSocketServer)wsUpgrade(w http.ResponseWriter, r *http.Request){
	serveWs(ws.hub, w, r)
}

func (ws *WebSocketServer)bindAddr(addr string){
	ws.addr = addr
}

func (ws *WebSocketServer)Start(){
	go ws.hub.run()
	http.HandleFunc("ws", ws.wsUpgrade)
	err := http.ListenAndServe(ws.addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}

	log.Fatal("start success...")
}

