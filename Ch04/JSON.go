/* ASN.1
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	mdata, err := json.Marshal(13)
	checkError(err)

	var n []uint8
	err = json.Unmarshal(n, mdata)
	checkError(err)

	fmt.Println("After marshal/unmarshal: ", n)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
