package ws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//InitWebSocketService 初始化websocket服务
func InitWebSocketService() (err error) {

	http.HandleFunc("/ws", serverWs)
	http.HandleFunc("/push", push)
	http.HandleFunc("/", indexPage)

	return http.ListenAndServe(":8080", nil)
}

//客户连接服务器
func serverWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(w, "init ws err:%s\n", err.Error())
		return
	}

	fmt.Printf("one client connected!,%s\n", c.RemoteAddr().String())

	cli := clients.addConn("1", c)

	go func() {
		defer c.Close()
		for {
			select {
			case msg := <-cli.msgChan:
				if err := cli.conn.WriteJSON(msg); err != nil {
					fmt.Println(err.Error())
				}
			default:
			}
		}
	}()

}

func push(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "param parse err:%s", err.Error())
		return
	}

	msg := r.FormValue("msg")

	pushData := pushData{
		PushTime: time.Now().Unix(),
		Msg:      msg,
		From:     "1",
	}

	var jsonData []byte
	var err error

	if jsonData, err = json.Marshal(pushData); err != nil {
		fmt.Printf("json encode err:%s", err.Error())
	}

	for _, c := range clients {
		c.msgChan <- string(jsonData)
	}
	// id := r.Form.Get("id")

	w.WriteHeader(200)

	// header := http.Header{
	// 	"content-Type": "application/json",
	// }
	// fmt.Fprint(w, "send success!")
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("./client/client.html", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Fprintf(w, "open file err:%s\n", err.Error())
		return
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "read file err:%s\n", err.Error())
		return
	}

	fmt.Fprintln(w, string(bytes))
}
