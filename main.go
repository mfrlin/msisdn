package main

import (
	"msisdn/parser"
	"fmt"
)

func main() {
	fmt.Print(parser.Parse_msisdn("38631661482"))
}
