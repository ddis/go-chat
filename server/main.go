package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net"
)

var connections []net.Conn

func main() {

	listener, _ := net.Listen("tcp", "192.168.31.27:8080")
	fmt.Println("Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go addConntection(conn)
		go handleClient(conn)
	}
}

func addConntection(conn net.Conn) {
	connections = append(connections, conn)

	fmt.Println("Added to list:" + conn.RemoteAddr().String())
}

func handleClient(conn net.Conn) {
	defer conn.Close() // закрываем сокет при выходе из функции
	fmt.Println("New connection")

	conn.Write([]byte("Connected\n"))

	buf := make([]byte, 32) // буфер для чтения клиентских данных
	for {
		readLen, err := conn.Read(buf) // читаем из сокета
		fmt.Println(readLen, "1")
		if err != nil {
			fmt.Println(err)
			break
		}

		for _, data := range connections {

			if conn.RemoteAddr().String() == data.RemoteAddr().String() {
				continue
			}

			data.Write(buf[:readLen]) // пишем в сокет
		}
	}
}

func getHash(addr string) string {
	h := sha256.New()
	h.Write([]byte(addr))

	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
