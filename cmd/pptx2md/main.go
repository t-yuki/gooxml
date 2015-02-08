package main

import (
	"fmt"
	"log"
	"os"

	"github.com/t-yuki/gooxml/ooxml"
)

func main() {
	var doc *ooxml.PresentationDocument
	var err error
	if len(os.Args) > 1 {
		doc, err = ooxml.OpenPresentationDocument(os.Args[1])
	} else {
		doc, err = ooxml.ReadPresentationDocument(os.Stdin)
	}
	if err != nil || doc == nil {
		log.Fatal(err)
	}
	slides := doc.Slides()
	for _, slide := range slides {
		fmt.Println(slide.String())
	}
}
