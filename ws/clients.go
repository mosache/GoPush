package ws

import "github.com/gorilla/websocket"

//客户端结构体
type client struct {
	id      string          //客户端id
	conn    *websocket.Conn //服务器与客户端的连接
	msgChan chan string
}

type clientSlice []*client //连接集合

var (
	clients clientSlice
)

//添加连接
func (cs clientSlice) addConn(id string, c *websocket.Conn) *client {
	client := &client{
		id:      id,
		conn:    c,
		msgChan: make(chan string),
	}
	cs = append(cs, client)

	return client
}

func (cs clientSlice) get(index int) *client {
	return cs[0]
}
