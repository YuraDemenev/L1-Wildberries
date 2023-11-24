package main

import "fmt"

type Sender interface {
	sendMessage()
}
type Message struct {
}

func (m *Message) EncryptionMessage() {
	fmt.Println("Message was encrypt")
}

type AdapterMessage struct {
	*Message
}

func (a *AdapterMessage) sendMessage() {
	a.EncryptionMessage()
	fmt.Println("Send message")
}

func Messanger(send Sender) {
	send.sendMessage()
}

func main() {
	message := &Message{}
	adapter := &AdapterMessage{message}
	Messanger(adapter)
}
