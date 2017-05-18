/* HTTPSFileServer
 */

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// deliver files from the directory /var/www 
	fileServer := http.FileServer(http.Dir("/var/www"))

	// register the handler and deliver requests to it
	err := http.ListenAndServeTLS(":8000", "jan.newmarch.name.pem",
		"private.pem", fileServer)
	//err := http.ListenAndServeTLS(":8000", "cert.pem",
	//	"key.pem", fileServer)
	checkError(err)
	// That's it!
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
