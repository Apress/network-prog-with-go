/* Marshal
 */

package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string "attr"
	Address string "chardata"
}

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan"},
			Email{Kind: "work", Address: "jan"}}}

	buff, _ := xml.Marshal(person)
	fmt.Println(string(buff))
}
