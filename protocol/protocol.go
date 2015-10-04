package protocol

import (
	"fmt"
)

type Protocol struct {
	Action      uint8
	RoomNumber  uint64
	UserIdSize  uint8
	UserId      string
	ContentSize uint
	Content     string
}

func (protocol *Protocol) Encode() (result []byte) {
	result = append(result, protocol.Action)
	result = append(result, byte(protocol.RoomNumber))
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
	protocol.RoomNumber = uint64(values[1])
	protocol.UserIdSize = values[2]
	protocol.ContentSize = uint(values[3])
	protocol.UserId = string(values[4 : 4+protocol.UserIdSize])
	protocol.Content = string(values[5+uint(protocol.UserIdSize) : 5+uint(protocol.UserIdSize)+protocol.ContentSize+1])

}
