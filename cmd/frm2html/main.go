package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xyproto/simpleform"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Please provide a filename as the first argument")
		os.Exit(1)
	}

	filename := os.Args[1]

	loginData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	htmlString, err := simpleform.HTML(string(loginData), true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(htmlString)
}
