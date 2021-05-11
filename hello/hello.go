package hello

import (
	"fmt"
	"rsc.io/quote"
)

func Say() {
	// Way 1
	message := NewMessage()
	message.setData(" | 001")

	// Way 2
	var message1 Message
	SetData(&message1, " | 002")

	fmt.Println(quote.Go() + message.getData() + message1.getData())
}
