package main

import (
	"fmt"
)

type Protocol struct {
	id     uint
	userid string
	action uint8
	msg    string
}

func main() {
	protocol := Protocol{
		id:     1,
		userid: "test",
		action: 10,
		msg:    "hello world",
	}

	msg := []byte(protocol.msg)
	fmt.Println("msg to byte : ", msg)
	fmt.Println("msg byte size : ", len(msg))
	fmt.Println("byte to string : ", string(msg))

}
