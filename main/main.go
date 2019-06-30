package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/xiahongze/id3conv"
)

func main() {

	flag.Parse()
	tails := flag.Args()

	if total := len(tails); total == 0 {
		fmt.Printf(`fix music file tagging encoding problems
id3conv usage:
	id3conv [files...]
At least one file should be given. Currently it only supports convert from GBK to UTF8.
`)
	} else {
		for _, f := range tails {
			log.Printf("process file \"%s\"", f)
			id3conv.Convert(f)
		}
		log.Printf("have processed %d files", total)
	}
}
