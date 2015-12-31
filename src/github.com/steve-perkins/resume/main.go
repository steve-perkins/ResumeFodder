package main

import (
	"flag"
	"fmt"
)

var format string

func main() {
	flag.StringVar(&format, "f", "default", "File format")
	flag.Parse()
	fmt.Printf("hello %s!\n", format)
//	flag.PrintDefaults()


	// Init a resume data file, in XML or JSON format

	// Convert to/from XML and JSON format

	// Generate resume output from data file
}
