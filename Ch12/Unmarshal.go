/* Unmarshal
 */

package main

import (
	"encoding/xml"
	"fmt"
	"os"
	//"strings"
)

type Person struct {
	XMLName Name    `xml:"person"`
	Name    Name    `xml:"name"`
	Email   []Email `xml:"email"`
}

type Name struct {
	Family   string `xml:"family"`
	Personal string `xml:"personal"`
}

type Email struct {
	Type    string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

func main() {
	str := `<?xml version="1.0" encoding="utf-8"?>
<person>
  <name>
    <family> Newmarch </family>
    <personal> Jan </personal>
  </name>
  <email type="personal">
    jan@newmarch.name
  </email>
  <email type="work">
    j.newmarch@boxhill.edu.au
  </email>
</person>`

	var person Person

	err := xml.Unmarshal([]byte(str), &person)
	checkError(err)

	// now use the person structure e.g.
	fmt.Println("Family name: \"" + person.Name.Family + "\"")
	fmt.Println("Second email address: \"" + person.Email[1].Address + "\"")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
