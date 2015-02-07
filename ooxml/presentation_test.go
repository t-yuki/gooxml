package ooxml

import (
	"strings"
	"testing"
)

func TestOpenPresentationDocument(t *testing.T) {
	doc, err := OpenPresentationDocument("../testdata/06.Hook into SharePoint APIs with Android.pptx")
	if err != nil {
		t.Fatal(err)
	}
	if doc == nil {
		t.Fatal(err)
	}
	slides := doc.Slides()
	if len(slides) == 0 {
		t.Fatal(doc)
	}
	if txt := doc.String(); strings.TrimSpace(txt) == "" {
		t.Fatal(doc)
	}
	t.Log(doc.String())
}
