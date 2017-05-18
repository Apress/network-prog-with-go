/* TemperatureServer
 */
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var ROOT_DIR = "/home/httpd/html/golang-hidden/websockets"

func GetTemp(ws *websocket.Conn) {
	for {

		msg, _ := exec.Command("sensors").Output()
		fmt.Println("Sending to client: " + string(msg[:]))
		err := websocket.Message.Send(ws, string(msg[:]))
		if err != nil {
			fmt.Println("Can't send")
			break
		}
		time.Sleep(2 * 1000 * 1000 * 1000)

		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir(ROOT_DIR))
	http.Handle("/GetTemp", websocket.Handler(GetTemp))
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
