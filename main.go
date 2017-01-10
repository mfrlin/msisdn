package main

import (
	"msisdn/parser"
	"fmt"
)

func main() {
	fmt.Print(parser.ParseMsisdn("38631123123"))
}
