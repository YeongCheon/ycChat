package main

import (
	"bufio"
	"chatting/protocol"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")

	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	var i uint8 = 11

	for {
		protocol := protocol.Protocol{
			Action:      1,
			UserIdSize:  4,
			UserId:      "test",
			ContentSize: 5,
			Content:     "hello",
		}

		writer := bufio.NewWriter(conn)
		writer.Write(protocol.Encode())
		writer.Flush()

		time.Sleep(2000 * time.Millisecond)
		i++
	}
}
