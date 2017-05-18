/* StrLen
 */

package main

import "fmt"

func main() {
	str := "百度一下，你就知道"

	fmt.Println("String length", len([]rune(str)))
	fmt.Println("Byte length", len(str))

	bytes := []byte(str)
	str2 := string(bytes[0 : len(bytes)-4])
	fmt.Println(len(str2), len([]rune(str2)))
	fmt.Println("\"" + str2 + "\"")

	for n, v := range []byte(str2) {
		fmt.Print(" (", n, " ", v, ")")
	}
	fmt.Println()
	for n, v := range str2 {
		fmt.Print(" (", n, " ", v, ")")
	}
	fmt.Println()
	for n, v := range []rune(str2) {
		fmt.Print(" (", n, " ", v, ")")
	}
	fmt.Println()

	str = "hello"
	fmt.Println("Hello String length", len([]rune(str)))
	fmt.Println("Hello Byte length", len(str))

}
