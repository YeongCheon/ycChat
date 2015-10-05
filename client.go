package main

import (
	"bufio"
	"chatting/protocol"
	"fmt"
	"net"
	"os"
)

func main() {
	const serverIP string = "211.208.161.151:6666"
	conn, err := net.Dial("tcp", serverIP)

	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Println("connected")

	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("tell me your nicname : ")
	nicname, _ := consoleReader.ReadString('\n')
	nicnameSize := uint32(len(nicname))

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
		UserIdSize:  nicnameSize,
		UserId:      nicname,
		ContentSize: 0,
	}

	writer := bufio.NewWriter(conn)
	writer.Write(joinProtocol.Encode())
	writer.Flush()

	for {
		message, _ := consoleReader.ReadString('\n')
		messageSize := uint(len(message))
		protocol := protocol.Protocol{
			Action:      3,
			RoomNumber:  1,
			UserIdSize:  nicnameSize,
			UserId:      nicname,
			ContentSize: messageSize,
			Content:     message,
		}

		writer.Write(protocol.Encode())
		writer.Flush()
	}
}
