package main

import (
	"bufio"
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
		fmt.Println("welcome")
		reader := bufio.NewReader(conn)
		go func(reader *bufio.Reader) {
			for {
				message, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(message)
			}
		}(reader)

	}
}
