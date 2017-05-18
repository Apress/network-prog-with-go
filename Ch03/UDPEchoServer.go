/* UDPEchoServer
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := "0.0.0.0:1200"
	udpAddr, err := net.ResolveUDPAddr("up4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {

	var buf [512]byte

	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println(string(buf[0:n]))

	daytime := time.Now().String()
	fmt.Println(daytime)

	conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
