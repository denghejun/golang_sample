package hello

import (
	"fmt"
	"rsc.io/quote"
)

func Say() {
	message := NewMessage()
	message.setData(" : 001")
	fmt.Println(quote.Go() + message.getData())
}
