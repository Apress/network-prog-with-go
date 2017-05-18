/* Aes
 */

package main

import (
	"bytes"
	"crypto/aes"
	"fmt"
)

func main() {
	key := []byte("my key, len 16 b")
	cipher, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	src := []byte("hello 16 b block")

	var enc [16]byte
	cipher.Encrypt(enc[0:], src)

	var decrypt [16]byte
	cipher.Decrypt(decrypt[0:], enc[0:])
	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:])
	fmt.Println(string(result.Bytes()))
}
