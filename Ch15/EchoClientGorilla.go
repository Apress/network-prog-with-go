/* EchoClientGorilla
 */
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	header := make(http.Header)
	header.Add("Origin", "http://localhost:12345")
	conn, _, err := websocket.DefaultDialer.Dial(service, header)
	checkError(err)

	for {
		//fmt.Println(`Readfnig msg`)
		_, reply, err := conn.ReadMessage()
		//fmt.Println(`Read msg`, err)
		if err != nil {
			// fmt.Println(err)
			if err == io.EOF {
				// graceful shutdown by server
				fmt.Println(`EOF from server`)
				break
			}
			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				fmt.Println(`Close from server`)
				break
			}
			fmt.Println("Couldn't receive msg " + err.Error())
			break
		}
		//checkError(err)
		fmt.Println("Received from server: " + string(reply[:]))

		// return the msg
		err = conn.WriteMessage(websocket.TextMessage, reply)
		if err != nil {
			fmt.Println("Couldn't return msg")
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
