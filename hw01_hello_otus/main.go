package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	inStr := "Hello, OTUS!"
	fmt.Print(stringutil.Reverse(inStr))
}
