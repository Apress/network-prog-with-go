package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"os"
)

//import ("fmt"; "rpc"; "os"; "net"; "log"; "http")

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith struct {
	Mult func(args *Args, reply *int) error
}

/*
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errorString("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
*/

func main() {

	mult := Arith{Mult: func(args *Args, reply *int) error {
		*reply = args.A * args.B
		return nil
	}}
	rpc.Register(mult)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
