package protocol

import (
	"fmt"
)

type Protocol struct {
	Action      uint8
	UserIdSize  uint8
	UserId      string
	ContentSize uint
	Content     string
}

func (protocol *Protocol) Encode() (result []byte) {
	result = append(result, protocol.Action)
	result = append(result, protocol.UserIdSize)
	result = append(result, byte(protocol.ContentSize))
	result = append(result, []byte(protocol.UserId)...)
	result = append(result, []byte(protocol.Content)...)
	return
}

func (protocol *Protocol) Decode(values []byte) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in protocol.Decode", r)
		}
	}()

	protocol.Action = values[0]
	protocol.UserIdSize = values[1]
	protocol.ContentSize = uint(values[2])
	protocol.UserId = string(values[3 : 3+protocol.UserIdSize])
	protocol.Content = string(values[3+uint(protocol.UserIdSize) : 3+uint(protocol.UserIdSize)+protocol.ContentSize+1])

}
