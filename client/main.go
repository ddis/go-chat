package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	serv := "192.168.31.27:8080"     // берем адрес сервера из аргументов командной строки
	conn, _ := net.Dial("tcp", serv) // открываем TCP-соединение к серверу
	go copyTo(os.Stdout, conn)       // читаем из сокета в stdout
	copyTo(conn, os.Stdin)           // пишем в сокет из stdin
}

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
