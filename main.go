package main

import (
	"fmt"
	"os"
)

func main() {
	cmdArgs := os.Args
	urlArgs := cmdArgs[1:]
	if len(urlArgs) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(urlArgs) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of %s\n", urlArgs[0])
}	