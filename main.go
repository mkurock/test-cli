package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Server struct {
	clients []net.Conn
	c chan string
}
var s = Server{}

func main() {
	fmt.Println("starting")
	s.c = make(chan string)
	l, _ := net.Listen("tcp", ":8080")
	go handleMsgs(s.c)
	for {
		con, _ := l.Accept()
		fmt.Println(con.RemoteAddr().String())
		go handleClient(con, s.c)
		s.clients = append(s.clients, con)
	}
}

func handleMsgs(ch <-chan string) {
	for msg := range ch {
		for _, client := range s.clients {
			client.Write([]byte(msg + "\n"))
		}
	}
}

func handleClient(c net.Conn, ch chan<- string) {
	for {
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\n\r")
		ch<-msg
		fmt.Println(msg)
	}
}