/* MD5Hash
 */

package main

import (
	"crypto/md5"
	"crypto/hmac"
	"fmt"
)

func main() {
	hash := hmac.New(md5.New, []byte("secret"))
	bytes := []byte("hello")
	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
}
