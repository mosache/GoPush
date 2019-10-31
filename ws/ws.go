package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

//InitWebSocketService 初始化websocket服务
func InitWebSocketService() (err error) {

	http.HandleFunc("/ws", serverWs)
	http.HandleFunc("/push", push)

	return http.ListenAndServe(":8080", nil)
}

//客户连接服务器
func serverWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(w, "init ws err:%s", err.Error())
		return
	}

	defer c.Close()

	cli := clients.addConn("1", c)

	select {
	case msg := <-cli.msgChan:
		fmt.Fprintf(w, "message:%s", msg)
	default:
	}
}

func push(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "param parse err:%s", err.Error())
		return
	}

	// id := r.Form.Get("id")

	c := clients.get(0)

	c.msgChan <- "hehe"

	fmt.Fprint(w, "send success!")
}
