/**
 * Base64
 */

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	eightBitData := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	enc := base64.StdEncoding.EncodeToString(eightBitData)
	dec, _ := base64.StdEncoding.DecodeString(enc)

	fmt.Println("Original data ", eightBitData)
	fmt.Println("Encoded string ", enc)
	fmt.Println("Decoded data ", dec)
}
