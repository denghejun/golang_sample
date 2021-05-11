package hello

import (
	"fmt"
	"rsc.io/quote"
)

func Say() {
	// Way 1
	message1 := NewMessage()
	message1.setData(" | 001")

	// Way 2
	var message2 Message
	SetData(&message2, " | 002")

	// Way 3
	var message3 Message
	message3.data = " | 003"

	fmt.Println(quote.Go() + message1.getData() + message2.getData() + message3.getData())
}
