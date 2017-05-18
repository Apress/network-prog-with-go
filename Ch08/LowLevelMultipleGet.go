/* LowLevelGet
 */

package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}
	url_, err := url.Parse(os.Args[1])
	checkError(err)

	for n := 0; n < 1000; n++ {
		// build a TCP connection to the proxy first
		host := url_.Host
		conn, err := net.Dial("tcp", host)
		checkError(err)

		// then wrap an HTTP client connection around it
		clientConn := httputil.NewClientConn(conn, nil)
		if clientConn == nil {
			fmt.Println("Can't build connection")
			os.Exit(1)
		}

		request := http.Request{Method: "GET", URL: url_}
		dump, _ := httputil.DumpRequest(&request, false)
		fmt.Println(string(dump))

		// send the request
		err = clientConn.Write(&request)
		checkError(err)

		// and get the response
		response, err := clientConn.Read(&request)
		checkError(err)

		var buf [512]byte
		reader := response.Body
		for {
			n, err := reader.Read(buf[0:])
			if err != nil {
				break
			}
			fmt.Print(string(buf[0:n]))
		}
		clientConn.Close()
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
