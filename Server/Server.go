package Server

import (
	"RsaOracle/Handler"
	"fmt"
	"net"
)

func TcpServer() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	defer l.Close()
	if err != nil {
		fmt.Println("Error while Listening")
		return
	}
	fmt.Println("Listening on localhost:8080")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error while Accepting")
			return
		}
		go Handler.TcpHandle(conn)
	}
}
