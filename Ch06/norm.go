package main

import (
	"fmt"
	"golang.org/x/text/unicode/norm"
)

func main() {
	str1 := "\u04d6"
	str2 := "\u0415\u0306"
	norm_str2 := norm.NFC.String(str2)
	bytes1 := []byte(str1)
	bytes2 := []byte(str2)
	norm_bytes2 := []byte(norm_str2)

	fmt.Println("Single char ", str1, " bytes ", bytes1)
	fmt.Println("Composed char ", str2, " bytes ", bytes2)
	fmt.Println("Normalised char", norm_str2, " bytes ", norm_bytes2)
	//fmt.Println(str1, bytes1, str2, bytes2)
}
