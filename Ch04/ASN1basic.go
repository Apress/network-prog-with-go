/* ASN.1 Basic
 */

package main

import (
	"encoding/asn1"
	"fmt"
	"os"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println("Before marshalling: ", t.String())

	mdata, err := asn1.Marshal(t)
	checkError(err)
	fmt.Println("Marshalled ok")

	var newtime = new(time.Time)
	_, err1 := asn1.Unmarshal(mdata, newtime)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", newtime.String())

	s := "hello \u00bc"
	fmt.Println("Before marshalling: ", s)

	mdata2, err := asn1.Marshal(s)
	checkError(err)
	fmt.Println("Marshalled ok")

	var newstr string
	_, err2 := asn1.Unmarshal(mdata2, &newstr)
	checkError(err2)

	fmt.Println("After marshal/unmarshal: ", newstr)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
