/* ASN1. Time
 */

package main

import (
	"encoding/asn1"
	"fmt"
	"os"
	"time"
)

// import "/.myTypes"

type MyInt int

func main() {
	t := time.Now()
	fmt.Println("Before marshalling: ", t.String())

	mdata, err := asn1.MarshalToMemory(t)
	checkError(err)

	var newtime time.Time
	_, err1 := asn1.Unmarshal(&newtime, mdata)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", newtime.String())
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
