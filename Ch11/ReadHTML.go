/* Read HTML
 */

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "file")
		os.Exit(1)
	}
	file := os.Args[1]
	bytes, err := ioutil.ReadFile(file)
	checkError(err)
	r := strings.NewReader(string(bytes))

	z := html.NewTokenizer(r)

	depth := 0
	for {
		tt := z.Next()

		for n := 0; n < depth; n++ {
			fmt.Print(" ")
		}

		switch tt {
		case html.ErrorToken:
			fmt.Println("Error ", z.Err().Error())
			os.Exit(0)
		case html.TextToken:
			fmt.Println("Text: \"" + z.Token().String() + "\"")
		case html.StartTagToken, html.EndTagToken:
			fmt.Println("Tag: \"" + z.Token().String() + "\"")
			if tt == html.StartTagToken {
				depth++
			} else {
				depth--
			}
		}
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
