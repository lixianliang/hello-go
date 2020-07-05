package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	file = flag.String("file", "", "input file")
)

func main() {
	flag.Parse()

	if *file == "" {
		fmt.Println("input file")
		return
	}

	fname := *file
	body, err := ioutil.ReadFile(fname)
	x := body[:500]
	if err != nil {
		panic(err)
	}
	fmt.Println(http.DetectContentType(x))
}
