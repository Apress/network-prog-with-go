/* StrLen
 */

package main

import "fmt"

var charmap = [256]int{0, 1, 2, 3, 4, 5}

var isoToUnicodeMap = map[uint8]int{
	0xc7: 0x12e,
	0xc8: 0x10c,
	0xca: 0x118,
}

var unicodeToISOMap = map[int]uint8{
	0x12e: 0xc7,
	0x10c: 0xc8,
	0x118: 0xca,
}

func main() {
	// str := "ĮČĘ"
	str := "\u012e\u010c\u0118"
	fmt.Println(str)

	bytes := unicodeStrToISO(str)

	// bytes is now the ISO 8859-1 encoding of the str
	for _, v := range bytes {
		fmt.Printf("  %x", v)
	}
	fmt.Println()

	// now turn it back into a UTF-8 string

	newStr := isoBytesToUnicode(bytes)
	fmt.Println(newStr)
}

/* Turn a UTF-8 string into an ISO 8859 encoded byte array
 */
func unicodeStrToISO(str string) []byte {
	codePoints := []int(str)
	// coedPoints contains unicode code points
	bytes := make([]byte, len(codePoints))
	for n, v := range codePoints {
		iso, ok := unicodeToISOMap[v]
		if !ok {
			iso = uint8(v)
		}
		bytes[n] = iso
	}
	return bytes
}

func isoBytesToUnicode(bytes []byte) string {
	newInts := make([]int, len(bytes))
	for n, v := range bytes {
		unicode, ok := isoToUnicodeMap[v]
		if !ok {
			unicode = int(v)
		}
		newInts[n] = unicode
	}
	newStr := string(newInts)
	return newStr
}
