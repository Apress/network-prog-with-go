/*
 * This program serves a file in preformatted, code layout
 * Useful for showing program text, properly escaping special
 * characters like '<', '>' and '&'
 */

package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", escapeString)

	err := http.ListenAndServe(":8080", nil)
	checkError(err)
}

func escapeString(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	bytes, err := ioutil.ReadFile("." + req.URL.Path)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	escapedStr := html.EscapeString(string(bytes))
	htmlText := "<html><body><pre><code>" +
		escapedStr +
		" </code></pre></body></html>"
	rw.Write([]byte(htmlText))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())
		os.Exit(1)
	}
}
