/* PersistentClient
 */

package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}
	url, err := url.Parse(os.Args[1])
	checkError(err)

	// build a TCP connection to the proxy first
	host := url.Host
	conn, err := net.Dial("tcp", host)
	checkError(err)

	// then wrap an HTTP client connection around it
	clientConn := httputil.NewClientConn(conn, nil)
	if clientConn == nil {
		fmt.Println("Can't build connection")
		os.Exit(1)
	}

	request := http.Request{Method: "GET", URL: url,
		Close: false, Proto: "HTTP/1.1",
		ProtoMajor: 1, ProtoMinor: 1}
	dump, _ := httputil.DumpRequest(&request, true)
	fmt.Println(string(dump))

	for n := 0; n < 1000; n++ {
		// send the request
		err = clientConn.Write(&request)
		if err == httputil.ErrPersistEOF {
			fmt.Println("Persist closed on write")
			break
		}
		checkError(err)

		// and get the response
		response, err := clientConn.Read(&request)
		if err == httputil.ErrPersistEOF {
			fmt.Println("Persist closed on read")
			break
		}

		checkError(err)

		var buf [512]byte
		reader := response.Body
		for {
			n, err := reader.Read(buf[0:])
			if err != nil {
				fmt.Println("No body")
				break
			}
			fmt.Print(string(buf[0:n]))
		}
		time.Sleep(1e9)
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
