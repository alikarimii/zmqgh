package value

func BuildMessage(b []byte) (Message, error) {
	// @TODO validate byte
	return Message(b), nil
}

func RebuildMessage(q string) Message {
	return Message(q)
}

type Message []byte

func (Message Message) String() string {
	return string(Message)
}
func (Message Message) Equals(other Message) bool {
	return Message.String() == other.String()
}
