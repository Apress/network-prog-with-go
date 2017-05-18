/* TLSEchoClient
 */
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	// Load the PEM self-signed certificate
	certPemFile, err := os.Open("jan.newmarch.name.pem")
	checkError(err)
	pemBytes := make([]byte, 1000) // bigger than the file
	_, err = certPemFile.Read(pemBytes)
	checkError(err)
	certPemFile.Close()

	// Create a new certificate pool
	certPool := x509.NewCertPool()
	// and add our certificate
	ok := certPool.AppendCertsFromPEM(pemBytes)
	if !ok {
		fmt.Println("PEM read failed")
	} else {
		fmt.Println("PEM read ok")
	}

	// Dial, using a config with root cert set to ours
	conn, err := tls.Dial("tcp", service, &tls.Config{RootCAs: certPool})
	checkError(err)

	// Now write and read lots
	for n := 0; n < 10; n++ {
		fmt.Println("Writing...")
		conn.Write([]byte("Hello " + string(n+48)))

		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError(err)

		fmt.Println(string(buf[0:n]))
	}
	conn.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
