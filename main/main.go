package main

import (
	"websocket/server"
)

func main(){
	wsSvr := server.NewWebSocketServer()
	wsSvr.Start()
}