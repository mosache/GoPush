package ws

type pushData struct {
	PushTime int64  `json:"pushTime"`
	Msg      string `json:"message"`
	From     string `json:"from"`
}
