package main

import (
	"bufio"
	"chatting/protocol"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, _ := listener.Accept()

		userReader := bufio.NewReader(conn)

		go func(reader *bufio.Reader) { //user join
			fmt.Println("welcome")
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
						fmt.Println("read complete")
						break
					}

				}

				packet := protocol.Protocol{}
				packet.Decode(values)

				fmt.Println(packet.Action)
				fmt.Println(packet.Content)
				fmt.Println(packet.UserId)
				fmt.Println(packet.ContentSize)
				fmt.Println(packet.Content)

			} //end for
		}(userReader)

	}
}
