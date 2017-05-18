/* ASN.1 EchoServer
 */
package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {

	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr(service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now()
		// Dereference the pointer from LocalTime.
		// Ignore return network errors.
		asn1.Marshal(conn, *daytime)
		conn.Close() // we're finished
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
