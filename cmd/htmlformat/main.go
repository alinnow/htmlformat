package main

import (
	"flag"
	"log"
	"os"

	"alin.ovh/htmlformat"
)

var (
	parseDocumentFlag = flag.Bool("document", false, "Set to true to parse a whole document")
	indentString      = flag.String("indent", "  ", "Set the indentation string")
)

func main() {
	flag.Parse()

	htmlformat.IndentString = *indentString

	var err error
	if *parseDocumentFlag {
		err = htmlformat.Document(os.Stdout, os.Stdin)
	} else {
		err = htmlformat.Fragment(os.Stdout, os.Stdin)
	}
	if err != nil {
		log.Fatalf("failed to format: %v", err)
	}
}
