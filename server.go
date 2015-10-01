package main

import (
	"bufio"
	"chatting/protocol"
	"fmt"
	"net"
	"strconv"
)

type Client struct {
	incoming chan protocol.Protocol
	outgoing chan protocol.Protocol
	reader   *bufio.Reader
	writer   *bufio.Writer
}

type Room struct {
	id      uint64
	clients []*Client
}

func (room *Room) addClient(client *Client) {
	room.clients = append(room.clients, client)
}

type Rooms struct {
	roomList []*Room
}

func (rooms *Rooms) FindRoom(findId uint64) bool {
	for _, room := range rooms.roomList {
		if room.id == findId {
			return true
		}
	}

	return false
}

func (rooms *Rooms) CreateRoom(roomNumber uint64) {
	room := Room{
		id: roomNumber,
	}

	rooms.roomList = append(rooms.roomList, &room)
}

func (rooms *Rooms) getRoom(roomNumber uint64) *Room {
	for _, room := range rooms.roomList {
		if room.id == roomNumber {
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
	fmt.Println("roomMember", room.clients)
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

		userReader := bufio.NewReader(conn)

		go func(reader *bufio.Reader) { //user join
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
					if roomNumber, _ := strconv.ParseUint(packet.Content, 10, 32); rooms.FindRoom(roomNumber) {
						rooms.JoinRoom(roomNumber, &Client{})
					} else {
						rooms.CreateRoom(roomNumber)
						rooms.JoinRoom(roomNumber, &Client{})
					}
				case 2: //leave room
					fmt.Println("leave room")
				case 3: //send message
					fmt.Println("send message")
				}
			} //end for
		}(userReader)

	}
}
