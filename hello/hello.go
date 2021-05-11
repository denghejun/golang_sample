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
	SetData(&message2, " | 002") // function always passed the VALUE not the REFERENCE/ADDR.

	// Way 3
	var message3 Message
	message3.data = " | 003"

	// Way 4
	var message4 Info // interface checking, if not implement the method listed in interface Info, there will be a type error.
	message4 = &Message{}
	message4.setData(" | 004")

	fmt.Println(quote.Go() + message1.getData() + message2.getData() + message3.getData() + message4.getData())
}
