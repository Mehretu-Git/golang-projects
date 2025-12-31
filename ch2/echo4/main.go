package main

import (
	"flag"
	"fmt"
	"strings"
)

// Initializing a new flag for -n and -s
var pointer_to_n = flag.Bool("n", false, "omit trailing newline")
var pointer_to_s = flag.String("s", " ", "separator")

func main() {
	//parsing command line flags
	flag.Parse()
	//joining the arguments with the specified separator
	fmt.Print(strings.Join(flag.Args(), *pointer_to_s))
	if !*pointer_to_n {
		//check if -n flag is not set, then print newline
		fmt.Println()
	}
}
