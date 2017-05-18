/* ClientGet
 */

package main

import (
	"fmt"
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

	client := &http.Client{}

	request, err := http.NewRequest("HEAD", url.String(), nil)

	// only accept UTF-8
	request.Header.Add("Accept-Charset", "utf-8;q=1, ISO-8859-1;q=0")
	checkError(err)

	response, err := client.Do(request)
	checkError(err)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	
	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))
	
	chSet := getCharset(response)
	if chSet != "utf-8" {
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
		return "utf-8"
	}
	idx := strings.Index(contentType, "charset=")
	if idx == -1 {
		// guess
		return "utf-8"
	}
	chSet := strings.Trim(contentType[idx+8:], " ")
	return strings.ToLower(chSet)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
