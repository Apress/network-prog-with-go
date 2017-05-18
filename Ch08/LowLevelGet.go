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
	"strings"
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

	// define the additional HTTP header fields
	//header := map[string] []string{"Accept-Charset": []string{"UTF-8;q=1", "ISO-8859-1;q=0"}}
	// and build the request
	//request := http.Request{Method: "GET", URL: url_, Proto: "HTTP/1.1", ProtoMajor: 1, ProtoMinor: 1} //, Header: header}
	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)

	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println(string(dump))

	// send the request
	err = clientConn.Write(request)
	checkError(err)

	// and get the response
	response, err := clientConn.Read(request)
	fmt.Println("Got response")
	checkError(err)

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	chSet := getCharset(response)
	fmt.Printf("got cahrset %s\n", chSet)
	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("got body")
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}

	os.Exit(0)
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		// guess
		return "UTF-8"
	}
	idx := strings.Index(contentType, "charset:")
	if idx == -1 {
		// guess
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], " ")
}

func checkError(err error) {
	if err != nil {
		if err == httputil.ErrPersistEOF {
			// ignore
			return
		}

		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
