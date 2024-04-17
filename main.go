package main

import (
	"RsaOracle/RsaCrypt"
	"RsaOracle/Server"
)

func main() {
	RsaCrypt.RsaInit()
	Server.TcpServer()
}
