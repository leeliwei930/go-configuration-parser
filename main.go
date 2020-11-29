package main

import (
	"flag"
	"fmt"
	"os"
	"techrino/utils"
)

var filename *string = flag.String("file", "example.txt", "The filename that you wish to parse it content to key and value")

func main() {
	flag.Parse()
	parser, err := utils.MakeParser(*filename)
	if err == nil {
		parser.Parse()
		parser.PrintContent()
		os.Exit(0)
	}

	fmt.Println(err)
	os.Exit(1)

}
