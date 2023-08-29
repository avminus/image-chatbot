package structures

var ImageIdToTypeMap = make(map[string]string)

type TextMessage struct {
	Content string `json:"message"`
}

type ImageMessage struct {
	Content []byte
}

var MessageList = make([]Message, 0)

type Messages struct {
	MessageList []Message `json:"messages"`
}

type Message struct {
	Role string      `json:"role"`
	Type string      `json:"type"`
	Id   string      `json:"id"`
	Data interface{} `json:"-"`
}

type MessageResponse struct {
	Status bool          `json:"status"`
	Data   []TextMessage `json:"data"`
}

type MetaMessageResponse struct {
	Status bool     `json:"status"`
	Data   Messages `json:"data"`
}
