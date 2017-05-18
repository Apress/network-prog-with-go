package main

import (
	"encoding/asn1"
	"fmt"
	//"time"
	"os"
)

// import "reflect"
// import "strings"
// import "/.myTypes"

type MyTime struct {
	Year                 int64 // 2006 is 2006
	Month, Day           int   // Jan-2 is 1, 2
	Hour, Minute, Second int   // 15:04:05 is 15, 4, 5.
	Weekday              int   // Sunday, Monday, ...
	ZoneOffset           int   // seconds east of UTC, e.g. -7*60 for -0700
	// Zone                 string // e.g., "MST"
}

type MyTime2 struct {
	Year                 int64  // 2006 is 2006
	Month, Day           int    // Jan-2 is 1, 2
	Hour, Minute, Second int    // 15:04:05 is 15, 4, 5.
	Weekday              int    // Sunday, Monday, ...
	ZoneOffset           int    // seconds east of UTC, e.g. -7*60 for -0700
	Zone                 string // e.g., "MST"
}

func main() {

	var t = new(MyTime)
	t.Year = 2010
	// fmt.Println("Before marshalling: ", t.String())

	mdata, err := asn1.Marshal(*t)
	checkError(err)
	fmt.Println("Marshalled ok")

	var newtime = new(MyTime2)
	_, err1 := asn1.Unmarshal(mdata, newtime)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", newtime.Year)

	/*
		var myNewtime MyTime
		_, err5 := asn1.Unmarshal(&myNewtime, mdata)
		checkError(err5)

		fmt.Println("After marshal/unmarshal: ", myNewtime.Year)




		m2, e6 := asn1.MarshalToMemory(13)
		checkError(e6)
		var n int
		_, e3 := asn1.Unmarshal(&n, m2)
		checkError(e3)

		fmt.Println("After marshal/unmarshal: ", n)
	*/
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

/*
func getType(t reflect.Type) string {
        switch t := t.(type) {
        case *reflect.BoolType:
                return "Boolean"
        case *reflect.IntType:
                return "Integer"
        case *reflect.StructType:
                return "Sequence"
        case *reflect.SliceType:
                if t.Elem().Kind() == reflect.Uint8 {
                        return "OctetString"
                }
                if strings.HasSuffix(t.Name(), "SET") {
                        return "Set"
                }
                return "Sequence"
        case *reflect.StringType:
                return "PrintableString"
	case *reflect.PtrType:
		return "Pointer"
	}

        return "error type"
}
*/

/*
func getUniversalType(t reflect.Type) (tagNumber int, isCompound, ok bool) {
        switch t {
        case reflect.objectIdentifierType:
                return tagOID, false, true
        case bitStringType:
                return tagBitString, false, true
        case timeType:
                return tagUTCTime, false, true
        case enumeratedType:
                return tagEnum, false, true
        }
        switch t := t.(type) {
        case *reflect.BoolType:
                return tagBoolean, false, true
        case *reflect.IntType:
                return tagInteger, false, true
        case *reflect.StructType:
                return tagSequence, true, true
        case *reflect.SliceType:
                if t.Elem().Kind() == reflect.Uint8 {
                        return tagOctetString, false, true
                }
                if strings.HasSuffix(t.Name(), "SET") {
                        return tagSet, true, true
                }
                return tagSequence, true, true
        case *reflect.StringType:
                return tagPrintableString, false, true
        }
        return 0, false, false
}
*/
