/* StrLen
 */

package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	str := "百度一下，你就知道"

	runess := utf16.Encode([]rune(str))
	ints := utf16.Decode(runes)
	fmt.Println(string(ints))

}
