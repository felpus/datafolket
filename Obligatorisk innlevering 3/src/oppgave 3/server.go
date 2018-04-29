package main

import (
	"net"
	"fmt"
	"os"
)

var (
	quote = "I want to die on Mars, just not on impact."
)
func main() {
	go TCP()
	go UDP()
	fmt.Println("Tjener starter")
}
func CheckError(err error) {
	if err != nil {
		fmt.Println("Feil: ", err)
		os.Exit(0)
	}
}
func TCP() {
	ln, err := net.Listen("tcp", ":17")
	CheckError(err)
	conn, err := ln.Accept()
	CheckError(err)
	defer conn.Close()
	for {
		_, err := conn.Write([]byte(quote))
		CheckError(err)
	}
}
func UDP() {
	address, err := net.ResolveUDPAddr("udp", ":17")
	CheckError(err)
	connection, err := net.ListenUDP("udp", address)
	CheckError(err)
	defer connection.Close()

	p := make([]byte, 128)

	for {
		_, addr, err := connection.ReadFromUDP(p)
		connection.WriteToUDP([]byte(quote), addr)
		CheckError(err)
	}
}
