package main

import (
	"bufio"
	"chatting/protocol"
	"fmt"
	"net"
)

type Client struct {
	Incoming chan protocol.Protocol
	Outgoing chan protocol.Protocol
	Reader   *bufio.Reader
	Writer   *bufio.Writer
}

type Room struct {
	Id      uint64
	Clients []*Client
}

func (room *Room) addClient(client *Client) {
	room.Clients = append(room.Clients, client)
}

func (room *Room) Broadcast(packet protocol.Protocol) {
	for _, client := range room.Clients {
		client.Writer.Write(packet.Encode())
		client.Writer.Flush()
	}
}

type Rooms struct {
	RoomList []*Room
}

func (rooms *Rooms) FindRoom(findId uint64) bool {
	for _, room := range rooms.RoomList {
		if room.Id == findId {
			return true
		}
	}

	return false
}

func (rooms *Rooms) CreateRoom(roomNumber uint64) {
	room := Room{
		Id: roomNumber,
	}

	rooms.RoomList = append(rooms.RoomList, &room)
}

func (rooms *Rooms) getRoom(roomNumber uint64) *Room {
	for _, room := range rooms.RoomList {
		if room.Id == roomNumber {
			fmt.Println("getRoom", room)
			return room
		}
	}

	return nil
}

func (rooms *Rooms) JoinRoom(roomNumber uint64, client *Client) bool {
	room := rooms.getRoom(roomNumber)
	if room == nil {
		fmt.Println("can not find room")
		return false
	}

	room.addClient(client)
	fmt.Println("roomMember", room.Clients)
	return true
}

func (rooms *Rooms) LeaveRoom(roomNumber uint64, client *Client) {

}

func main() {
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	rooms := Rooms{}

	for {
		conn, _ := listener.Accept()

		go func(conn net.Conn) { //user join
			reader := bufio.NewReader(conn)

			client := Client{
				Writer: bufio.NewWriter(conn),
			}
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

				switch packet.Action {
				case 1: //join room
					if roomNumber := packet.RoomNumber; rooms.FindRoom(roomNumber) {
						rooms.JoinRoom(roomNumber, &client)
					} else {
						rooms.CreateRoom(roomNumber)
						rooms.JoinRoom(roomNumber, &client)
					}
				case 2: //leave room
					fmt.Println("leave room")
				case 3: //broadcast
					fmt.Println("broadcast")
					room := rooms.getRoom(packet.RoomNumber)

					go room.Broadcast(packet)
				}
			} //end for
		}(conn)

	}
}
