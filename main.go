package main

import "GoPush/ws"

func main() {
	if err := ws.InitWebSocketService(); err != nil {
		panic(err.Error())
	}
}
