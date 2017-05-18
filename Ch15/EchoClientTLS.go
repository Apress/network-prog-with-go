/* EchoClientTLS
 */
package main

import (
	"fmt"
	"crypto/tls"
	"golang.org/x/net/websocket"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "wss://host:port")
		os.Exit(1)
	}

	config, err := websocket.NewConfig(os.Args[1], "http://localhost")
	checkError(err)
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	config.TlsConfig = tlsConfig
	
	conn, err := websocket.DialConfig(config)
	checkError(err)
	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				// graceful shutdown by server
				break
			}
			fmt.Println("Couldn't receive msg " + err.Error())
			break
		}
		fmt.Println("Received from server: " + msg)
		// return the msg
		err = websocket.Message.Send(conn, msg)
		if err != nil {
			fmt.Println("Coduln't return msg")
			break
		}
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
