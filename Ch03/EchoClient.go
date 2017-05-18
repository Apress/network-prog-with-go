/* EchoClient
 */
package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], " host:port\n")
		// fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError("Dial", err)

	for n := 0; n < 10; n++ {
		conn.Write([]byte("Hello " + string(n+48)))

		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError("Read", err)

		fmt.Println(string(buf[0:n]))
	}
	os.Exit(0)
}

func checkError(errStr string, err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, errStr, err.Error())
		//log.Exit(errStr+":",  err.Error())
		_, file, line, ok := runtime.Caller(1)
		if ok {
			fmt.Println("Line number", file, line)
		}
		os.Exit(1)
		/*
			fmt.Println("Fatal error ", err.Error())
			os.Exit(1)
		*/
	}
}
