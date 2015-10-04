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

	go func() {
		reader := bufio.NewReader(conn)
		for {
			var values []byte
			buf := make([]byte, 10)
			for {
				size, err := reader.Read(buf)

				if err != nil {
					fmt.Println(err)
					return
				}

				values = append(values, buf...)

				if size < len(buf) {
					break //packet read complete
				}
			}

			packet := protocol.Protocol{}
			packet.Decode(values)

			fmt.Println(packet)
		}
	}()

	joinProtocol := protocol.Protocol{
		Action:      1,
		RoomNumber:  1,
		UserIdSize:  4,
		UserId:      "test",
		ContentSize: 4,
		Content:     "join",
	}

	writer := bufio.NewWriter(conn)
	writer.Write(joinProtocol.Encode())
	writer.Flush()

	for {
		protocol := protocol.Protocol{
			Action:      3,
			RoomNumber:  1,
			UserIdSize:  4,
			UserId:      "test",
			ContentSize: 5,
			Content:     "prince",
		}

		writer.Write(protocol.Encode())
		writer.Flush()

		time.Sleep(2000 * time.Millisecond)
		i++
	}
}
