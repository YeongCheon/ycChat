package main

import (
	"bufio"
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

	for {
		writer := bufio.NewWriter(conn)
		writer.WriteString("hello")
		writer.WriteRune('\n')
		writer.Flush()
		fmt.Println("Done")
		time.Sleep(1000 * time.Millisecond)
	}
}
