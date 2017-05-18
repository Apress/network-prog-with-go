/* Persistent Server
 */

package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {

	// create a TCP server first
	listener, err := net.Listen("tcp", ":8000")
	checkError(err)

	// then wrap an HTTP client connection around it

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Accepted")
		go handleConnection(conn)
	}
	os.Exit(0)
}

func handleConnection(conn net.Conn) {
	serverConn := httputil.NewServerConn(conn, nil)
	if serverConn == nil {
		fmt.Println("Can't build connection")
		return
	}

	// timeout
	go func(serverConn *httputil.ServerConn) {
		time.Sleep(10e9) // 10 seconds
		serverConn.Close()
	}(serverConn)

	for {
		req, err := serverConn.Read()
		if err != nil {
			if err == httputil.ErrPersistEOF {
				// client has closed ok
				fmt.Println("Persistent conn closed on read")
				break
			}
			// unexpected error
			fmt.Println(err.Error())
			break
		}
		bytes, _ := httputil.DumpRequest(req, false)
		fmt.Println(string(bytes))

		// Prepare "no content" reponse
		response := http.Response{
			Status:     "204 No Content",
			StatusCode: 204,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Request:    req,
			Close:      false,
		}
		serverConn.Write(req, &response)
		fmt.Println("wrote response")
	}
	serverConn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
