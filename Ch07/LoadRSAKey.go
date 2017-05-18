/* GenRSAKeys
 */

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	reader := rand.Reader
	bitSize := 128
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	fmt.Println("Private key primes", key.P, key.Q)
	fmt.Println("Private key exponent", key.D)

	publicKey := key.PublicKey
	fmt.Println("Public key modulus", publicKey.N)
	fmt.Println("Public key exponent", publicKey.E)

	saveKey("private.key", key)
	saveKey("public.key", key.PublicKey)
}

func saveKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
