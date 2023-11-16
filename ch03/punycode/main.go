package main

import (
	"fmt"
	"golang.org/x/net/idna"
)

func main() {
	src := "お名前.com"
	ascii, err := idna.ToASCII(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s\n", src, ascii)
}
