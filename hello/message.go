package hello

type Message struct {
	data string
}

func NewMessage() *Message {
	return &Message{}
}

func SetData(msg *Message, data string) {
	msg.data = data
}

func (msg *Message) getData() string {
	return msg.data
}

func (msg *Message) setData(data string) {
	msg.data = data
}
