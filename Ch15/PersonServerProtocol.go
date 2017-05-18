/* PersonServerProtocol
 */
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"xmlcodec"
)

var protocolChosen string

type Person struct {
	Name   string
	Emails []string
}

func Echo(ws *websocket.Conn) {
	var person Person
	var err error

	protos := ws.Config().Protocol
	if len(protos) != 1 {
		os.Exit(1)
	}
	protocolChosen := protos[0]
	if protocolChosen == "json" {
		err = websocket.JSON.Receive(ws, &person)
	} else {
		err = xmlcodec.XMLCodec.Receive(ws, &person)
	}
	if err != nil {
		fmt.Println("Can't receive")
	} else {

		fmt.Println("Name: " + person.Name)
		for _, e := range person.Emails {
			fmt.Println("An email: " + e)
		}
	}
}

func Chooser(clientProtos []string) (string, error) {
	// See if any of the server's preferences are listed
	// in the client's offerings - server takes precedence
	acceptableProtos := []string{"xml", "json"}
	if len(clientProtos) > 1 {
		for _, p := range acceptableProtos {
			for _, q := range clientProtos {
				if p == q {
					return p, nil
				}
			}
		}
	}
	// no match
	return "", nil
}

func main() {
	handler := websocket.Handler(Echo)
	handler.ProtocolChooser = Chooser
	http.Handle("/", handler)
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
